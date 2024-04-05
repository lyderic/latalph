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
	echo "Tu fais de ton mieux."  | latalph

set shell := ["bash","-uc"]
# vim: ft=make
