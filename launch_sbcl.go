package main

import ()

func Launch_sbcl(config map[string]string, args []string) {
	bin_path := HomeDir()
	bin_path += "/" + "impls"
	bin_path += "/" + UnameM() // x86-64
	bin_path += "/" + UnameS() // darwin
	bin_path += "/" + "sbcl-bin"
	bin_path += "/" + config["sbcl-bin.version"]
	bin_path += "/" + "bin"
	bin_path += "/" + "sbcl"
	args[0] = bin_path
	Exec(bin_path, args, nil)
}
