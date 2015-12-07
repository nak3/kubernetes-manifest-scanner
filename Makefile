all build:
	hack/build.sh
.PHONY: all build


build_bashcomp:
	hack/bashcomp_build.sh
        # NOTE: At this moment, the cobra will make bash-completion which you don't expect due to this:
	# https://github.com/spf13/cobra/pull/205
        # Be careful when you auto genrate the bash completion
	bin/bash_comp_autogenerator
	echo "COMP_WORDBREAKS=\${COMP_WORDBREAKS//:}" >> completions/bash/kuberenetes-manifest-scanner
.PHONY: build_bashcomp


clean:
	rm -rf gopath
	rm -rf bin/*
	rm -rf _output
	#rm -rf Godeps
.PHONY: clean
