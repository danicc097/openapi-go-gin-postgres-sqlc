should simply move the api_{tag}.go file to handlers if api_{tag}.go does not exist there.
OPtions for services, etc.:
1. it will create a service file services/{tag}.go with an empty {Tag} struct if not exists as well, else wont compile. This will force us to use the service pattern, but only at the start, then if we remove it it wont be added again automagically, bad idea.
2. mustache doesnt create ``svc services.{tag}`` and its up to us to fill in (wont be removed/overwritten). More liberty, preferred.
