protogen:
	bazel query 'attr(generator_function, proto_gen, //...)' | xargs -n1 bazel run