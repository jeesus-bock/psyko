# psyko

A very experimental project investigating Go TCL integration.

Dynamic end-point generation based on available script files. The available
ones ought to be advertised on some general end-point maybe.

Yes, it doesn't make much sense, but I had just finished another project
and was bored and started playing around with Gin some more. The general
plan is that there is no plan.

Usin the bleediest edge Go, though should work with earlier versions too just fine.

### Front-end

Testing out AlpineJS again, I'm sort of tired of heavvy-duty fe systems and here
we don't have any sort of hard requirements so mightaswell experiment.

### Containerization

I want to test github actions for building aa docker image. Initial Dockerfile done, very
rudimentary and probably wrong. Also pushed this stuff to dockerhub. I wonder if there's
other as asinine stuff as this there :).
`docker build . -t jeesusbock/psyko:0.12`
`docker push jeesusbock/psyko:0.12`

### Other functionality
I also want persistence propably sqlite, but I want to make it accessible from
the tcl scripts and mainly the scripts.
