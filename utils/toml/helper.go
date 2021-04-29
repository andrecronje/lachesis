package toml

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/naoina/toml/ast"
)

var (
	// ErrorParamNotExists error for not exists params when read
	ErrorParamNotExists = errors.New("param not exists")
	// ErrorParamAlreadyExists error for already exists params when add
	ErrorParamAlreadyExists = errors.New("param already exists")
	// ErrorSectionNotExists error for not exists section when search
	ErrorSectionNotExists = errors.New("section not exists")
	// ErrorSectionAlreadyExists error for already exists section when add
	ErrorSectionAlreadyExists = errors.New("section already exists")
)

// Helper is helper for simple manipulate with parsed toml data
type Helper struct {
	table *ast.Table
}

// NewTomlHelper return new helper using *ast.Table created in toml library Parse method
func NewTomlHelper(t *ast.Table) *Helper {
	return &Helper{
		table: t,
	}
}

// GetTable return *ast.Table data
func (d *Helper) GetTable() *ast.Table {
	return d.table
}

// AddSection add section with name 'name' in toml data after section 'after'
func (d *Helper) AddSection(name, after string) error {
	if name == "" {
		return nil
	}

	_, err := d.FindSection(name)
	if err == nil {
		// If exists - return error
		return ErrorSectionAlreadyExists
	}

	path := strings.Split(name, ".")

	afterSection, err := d.FindSection(after)
	if err != nil {
		return err
	}

	pathStr := ""
	currentSection := d.table
	for _, n := range path {
		pathStr = pathStr + "/" + n

		var section *ast.Table
		sectionI, ok := currentSection.Fields[n]
		if ok {
			section, ok = sectionI.(*ast.Table)
			if !ok {
				return errors.New("wrong type of section: " + pathStr)
			}
		} else {
			section = &ast.Table{
				Position: ast.Position{
					Begin: afterSection.End() + 1,
					End:   0,
				},
				Line:   afterSection.Line + afterSection.End() - afterSection.Pos(),
				Name:   n,
				Fields: make(map[string]interface{}),
				Type:   ast.TableTypeNormal,
			}

		}
		currentSection.Fields[n] = section
		currentSection = section
	}

	return nil
}

// DeleteSection remove section named 'name' with subdata from toml data
func (d *Helper) DeleteSection(name string) error {
	// Find parent section and name for deletedName section
	path := strings.Split(name, ".")
	parentName := strings.Join(path[:len(path)-1], ".")
	deletedName := path[len(path)-1]

	parent, err := d.FindSection(parentName)
	if err != nil {
		return err
	}

	delete(parent.Fields, deletedName)

	return nil
}

// RenameSection change name for section
func (d *Helper) RenameSection(name, newName string) error {
	section, err := d.FindSection(name)
	if err != nil {
		return err
	}

	section.Name = newName
	delete(d.table.Fields, name)
	d.table.Fields[newName] = section

	return nil
}

// AddParam add param in section with set value
func (d *Helper) AddParam(name, sectionName string, value interface{}) error {
	_, sect, err := d.getKVData(name, sectionName)
	if err == nil {
		return ErrorParamAlreadyExists
	}
	if sect == nil {
		return err
	}

	kvData, err := d.setKVData(name, value)
	if err != nil {
		return err
	}

	sect.Fields[name] = kvData

	return nil
}

// DeleteParam remove param from section
func (d *Helper) DeleteParam(name, sectionName string) error {
	_, sect, err := d.getKVData(name, sectionName)
	if err != nil {
		return err
	}

	delete(sect.Fields, name)

	return nil
}

// RenameParam rename param
func (d *Helper) RenameParam(name, sectionName, newName string) error {
	param, sect, err := d.getKVData(name, sectionName)
	if err != nil {
		return err
	}

	delete(sect.Fields, name)
	param.Key = newName
	sect.Fields[newName] = param

	return nil
}

// SetParam set value for exists param
func (d *Helper) SetParam(name, sectionName string, value interface{}) error {
	param, _, err := d.getKVData(name, sectionName)
	if err != nil {
		return err
	}

	_, err = d.setKVData(name, value, param)
	if err != nil {
		return err
	}

	return nil
}

// GetParamString get value for string param
func (d *Helper) GetParamString(name, sectionName string) (string, error) {
	param, _, err := d.getKVData(name, sectionName)
	if err != nil {
		return "", err
	}

	pString, ok := param.Value.(*ast.String)
	if !ok {
		return "", errors.New("wrong type for string in param: " + sectionName + " / " + name)
	}

	return pString.Value, nil
}

// GetParamInt get value for int param
func (d *Helper) GetParamInt(name, sectionName string) (int64, error) {
	param, _, err := d.getKVData(name, sectionName)
	if err != nil {
		return -1, err
	}

	pInt, ok := param.Value.(*ast.Integer)
	if !ok {
		return -1, errors.New("wrong type for integer in param: " + sectionName + " / " + name)
	}

	return pInt.Int()
}

// GetParamFloat get value for float param
func (d *Helper) GetParamFloat(name, sectionName string) (float64, error) {
	param, _, err := d.getKVData(name, sectionName)
	if err != nil {
		return -1, err
	}

	pFloat, ok := param.Value.(*ast.Float)
	if !ok {
		return -1, errors.New("wrong type for integer in param: " + sectionName + " / " + name)
	}

	return pFloat.Float()
}

// GetParamBool get value for bool param
func (d *Helper) GetParamBool(name, sectionName string) (bool, error) {
	param, _, err := d.getKVData(name, sectionName)
	if err != nil {
		return false, err
	}

	pBool, ok := param.Value.(*ast.Boolean)
	if !ok {
		return false, errors.New("wrong type for integer in param: " + sectionName + " / " + name)
	}

	return pBool.Boolean()
}

// GetParamTime get value for Time param
func (d *Helper) GetParamTime(name, sectionName string) (time.Time, error) {
	param, _, err := d.getKVData(name, sectionName)
	if err != nil {
		return time.Now(), err
	}

	pTime, ok := param.Value.(*ast.Datetime)
	if !ok {
		return time.Now(), errors.New("wrong type for integer in param: " + sectionName + " / " + name)
	}

	return pTime.Time()
}

// FindSection return *ast.Table data for section or (nil, ErrorSectionNotExists) if section not exists
func (d *Helper) FindSection(name string) (*ast.Table, error) {
	path := strings.Split(name, ".")
	currentSection := d.table

	if name == "" {
		return currentSection, nil
	}

	pathStr := ""
	for _, n := range path {
		pathStr = pathStr + "/" + n

		sectionI, ok := currentSection.Fields[n]
		if !ok {
			return nil, ErrorSectionNotExists
		}

		currentSection, ok = sectionI.(*ast.Table)
		if !ok {
			return nil, errors.New("section has wrong type: " + pathStr)
		}
	}

	return currentSection, nil
}

func (d *Helper) getKVData(name, sectionName string) (*ast.KeyValue, *ast.Table, error) {
	sect, err := d.FindSection(sectionName)
	if err != nil {
		return nil, nil, err
	}
	if sect == nil {
		return nil, nil, ErrorSectionNotExists
	}

	paramI, ok := sect.Fields[name]
	if !ok {
		return nil, sect, ErrorParamNotExists
	}

	param, ok := paramI.(*ast.KeyValue)
	if !ok {
		return nil, sect, ErrorParamNotExists
	}

	return param, sect, nil
}

func (d *Helper) setKVData(name string, value interface{}, kvExists ...*ast.KeyValue) (*ast.KeyValue, error) {
	var kv *ast.KeyValue
	if len(kvExists) > 0 {
		kv = kvExists[0]
	} else {
		kv = &ast.KeyValue{
			Key: name,
		}
	}
	// TODO: check if Data should be quoted (as ast.String)
	switch v := value.(type) {
	case string:
		kv.Value = &ast.String{
			Value: v,
			Data:  []rune(`"` + v + `"`),
		}
	case int:
		s := strconv.FormatInt(int64(v), 10)
		kv.Value = &ast.Integer{
			Position: ast.Position{},
			Value:    s,
			Data:     []rune(s),
		}
	case int64:
		s := strconv.FormatInt(v, 10)
		kv.Value = &ast.Integer{
			Value: s,
			Data:  []rune(s),
		}
	case float64:
		s := strconv.FormatFloat(v, 'f', 16, 64)
		kv.Value = &ast.Float{
			Value: s,
			Data:  []rune(s),
		}
	case bool:
		s := strconv.FormatBool(v)
		kv.Value = &ast.Boolean{
			Value: s,
			Data:  []rune(s),
		}
	case time.Time:
		s := v.Format("2006-01-02T15:04:05.999999999Z07:00")
		kv.Value = &ast.Datetime{
			Value: s,
			Data:  []rune(s),
		}
	default:
		log.Panic("Bad type of value for config key " + name)
	}

	return kv, nil
}
