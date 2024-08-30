This testcase is to reproduce the inconsistency I am seeing when running `dagger call <fn>`.

To reproduce:

### Run Dev engine dagger

- git clone git@github.com:rajatjindal/dagger.git
- git checkout <branch name>
- ensure no dagger engine is running (checked via docker desktop and removed the running containers)
- start the dev engine by running `./hack/dev`


### Reproduce the issue

- git clone git@github.com:rajatjindal/dagger-views.git
- git checkout main
- Notice that `dagger.json` and `module-abc/dagger.json` points to engine version `v0.12.6` (the `strcase` lib change cutoff as of right now is v0.12.5)
- Run `/path/to/dagger/repo/hack/with-dev dagger develop`
- Notice that the `dagger.gen.go and internal/**` files are generated
- Run `/path/to/dagger/repo/hack/with-dev dagger call container-echo --skip-t-parse-xyz hello2 --string-arg hello` - this should work
- Run `/path/to/dagger/repo/hack/with-dev dagger call get-darwin-cli-version` - this should work as well
- Now change the engine version in `dagger.json` AND `module-abc/dagger.json` to `v0.12.5`
- Run `/path/to/dagger/repo/hack/with-dev dagger develop`
- Try running above command a few times, notice that `internal/dagger/dagger.gen.go` sometimes flips between having a function `GetKubectlCLIVersion` OR `GetKubectlCliversion`

My theory is that 

- when a call is made to server, the function `CachedSelect` may be called at different times in parallel and there may be a race condition of when setting ctx here. (or maybe picking up wrong cache?)
- It works with `v0.12.6` consistently because if the `strcase` is missing from the context, we default to the newer version.

Tracing the code on how the code generation is done, I reached `generator.go` under codegen, where we are rendering the templates, which makes use of schema and schemaVersion which are either queried or provided from command line. Looking at logs, it seems like it is being passed in a file on command line. 
