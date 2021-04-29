#!/usr/bin/env bash
cd $(dirname $0)
. ./_params.sh

set -e

echo -e "\nConfigure tx-storm:\n"

METRICS=--metrics

cat << HEADER > tx-storm.toml
ChainId = 4003

SendTrusted = false

URLs = [                                                                                                                                                                                                    
HEADER

for ((i=0;i<$N;i+=1))
do
  WSP=$(($WSP_BASE+$i))
  echo "\"ws://127.0.0.1:${WSP}\"", >> tx-storm.toml
done

cat << FOOTER >> tx-storm.toml
]

[Accs]
Count = ${TEST_ACCS_COUNT}
Offset = ${TEST_ACCS_START}
FOOTER

echo -e "\nStart tx-storm:\n"
(go run ../cmd/tx-storm \
    ${METRICS} --verbosity 3 \
    &> .txstorm.log)&

