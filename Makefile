
gen_default_client:
	go build -o bin/default_client .

gen_uniform_pgo_client:
	go build -o bin/uniform_pgo_client -pgo=profiles/uniform.pgo .

gen_read_heavy_pgo_client:
	go build -o bin/read_heavy_pgo_client -pgo=profiles/read_heavy.pgo .
