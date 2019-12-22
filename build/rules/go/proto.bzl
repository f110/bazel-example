load("@bazel_skylib//lib:shell.bzl", "shell")
load("@io_bazel_rules_go_compat//:compat.bzl", "get_proto")

def _proto_gen_impl(ctx):
    generated = ctx.attr.src[OutputGroupInfo].go_generated_srcs.to_list()
    substitutions = {
        "@@FROM@@": shell.quote(generated[0].path),
        "@@TO@@": shell.quote(ctx.attr.dir),
    }
    out = ctx.actions.declare_file(ctx.label.name + ".sh")
    ctx.actions.expand_template(
        template = ctx.file._template,
        output = out,
        substitutions = substitutions,
        is_executable = True,
    )
    runfiles = ctx.runfiles(files = [generated[0]])
    return [
        DefaultInfo(
            runfiles = runfiles,
            executable = out,
        ),
    ]

_proto_gen = rule(
    implementation = _proto_gen_impl,
    executable = True,
    attrs = {
        "dir": attr.string(),
        "src": attr.label(),
        "_template": attr.label(
            default = "//build/rules/go:move-into-workspace.bash",
            allow_single_file = True,
        ),
    },
)

def proto_gen(name, **kwargs):
    if not "dir" in kwargs:
        dir = native.package_name()
        kwargs["dir"] = dir

    _proto_gen(name = name, **kwargs)
