_listing:
	@printf "${BLU}{{justfile()}}${NOC}\n"
	@just --unsorted --list --list-heading='' --list-prefix=' • ' \
		| grep -v 'alias for'

install:
	go install -v

test: install
	echo "Jérusalem" | latalph
	echo "¡Un café, por favor!" | latalph
	echo "You are giving the maximum." | latalph
	echo "Dies kostet 30 Euros."  | latalph
	echo "La clé est '123password'."  | latalph

set shell := ["bash","-uc"]
