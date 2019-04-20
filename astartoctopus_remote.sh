AlphabetwebProject="/Users/vikingbays/golang/AlphabetwebProject/alphabetweb"

sample_octopus_frontend_service="/Users/vikingbays/golang/AlphabetwebProject/alphabet_samples/sample_octopus_frontend_service"

config_remote_url="http://localhost:10060/sysconfigservice/sysconfig/downzip/octopus/base_remote/"

go run $AlphabetwebProject/src/alphabet/cmd/abserver.go -start "$sample_octopus_frontend_service"  -config_url "$config_remote_url" -env_x1 y1 -env_x2 y2 -appsname "sample_octopus_frontend,sample_octopus_service"
