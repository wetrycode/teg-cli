# Golang-based tooling for tegenaria development  

 Tegenaria CLI is a scaffolding tool for [tegenaria](https://github.com/wetrycode/tegenaria).

 # Feature

 - New a tegenaria project base on template.  

 - Add a new spdier,pipeline or middlerware

 # Example

 - New a tegenaria project
 ```bash
 teg-cli new demo --output /data/work/wetrycode/demo --spider demo -m github.com/wetrycode/demo
 ```

 - Add a new spider

 ```bash
  teg-cli add spider --name Spider1 --filename spider1 -o ./demo/spiders
```

- Init an exist project  

```bash
teg-cli init --spider demo -m github.com/wetrycode/demo
```