def _example_config_test_impl(ctx):
    src = ctx.file.src
    kicker = ctx.actions.declare_file("%s_kicker.sh" % ctx.label.name)
    ctx.actions.expand_template(
        template = ctx.file._wrapper_template,
        output = kicker,
        substitutions = {
            "{executable_binary}": ctx.executable._test_suite.short_path,
            "{config_file}": src.short_path,
        },
     is_executable = True,
    )
    runfiles = ctx.runfiles(files = [kicker, ctx.executable._test_suite, src])
    return [DefaultInfo(executable = kicker, runfiles = runfiles)]

example_config_test = rule(
    implementation = _example_config_test_impl,
    test = True,
    attrs = {
        "src": attr.label(allow_single_file = True),
        "_test_suite": attr.label(
            default = Label("//test-suite/example-suite"),
            executable = True,
            cfg = "target",
        ),
        "_wrapper_template": attr.label(
            allow_single_file = True,
            default = "kicker.tpl",
        )
    },
)