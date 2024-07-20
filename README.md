# Go Blog
## Info
This is a hobby project which I started due to the fact that I am resolving programming problems on a daily basis and I wanted a sigle place where I could post my findings.

I am aware of multiple other webistes which provide this functionality, however I thought it would be fun to do this in go-htmx-mongo db.

Despite me recently studying bootstrap and tailwind I thought it was an absolute waste of time to tweak the design and went for a rather simplistic approach.

I ended up creating a cloud instance on [Hetzner](https://www.hetzner.com/), hosting via nginx and using [MongoDb Atlas](https://cloud.mongodb.com/) for storage.

I will possibly be making small tweaks in the future like hosting the db source elsewhere as the free tier on Atlas is only 512 megs, but the majority of the works seems to be completed.

The website is hosted at [https://blog.pragmatino.xyz](https://blog.pragmatino.xyz)

## Running locally
After git cloning or downloading open the directory and run 
> ./dev.sh
>
> which would execute the following commands
> 
```
Obtaining a css file from the tailwind library
echo "tw" && npx tailwindcss -i ./input.css -o ./components/styles/output.css &&

obtaining html files from the templ files
echo "templ" && templ generate && 

Running our go application
echo "go" && go run main.go

```
