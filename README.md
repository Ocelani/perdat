# ![perdat](internal/assets/perdat_light_2x.png)

A CLI tool that registers information about your daily life.
It collects the provided data inputed by the user and stores in a single sqlite file.
With that data, you can get some insights and know yourself, then, keep learning.

### Development stage

Know yourself and happy hacking!

## Requirements

**E.01** - Me, as a person who wants to know more about myself, I would like to register any kind of event or data in one private, secure and resilient platform. In order to keep collecting some infos and get some insights about my personality and my life.

**RF.01** - Should register one single event to the sqlite file with a given command, e.g.

```bash
pdt add <fact> [--date -d]
pdt add headache
pdt add created pull request -d yesterday
```

### Data Model

#### FACT

- ID : pk
- Name : string
- Day : time.Date (default time.Now)
- CratedAt : timestamp
- UpdatedAt : timestamp
