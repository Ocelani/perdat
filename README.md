# ![perdat](internal/assets/perdat_light_2x.png)

A CLI tool that registers information about your daily life.
It collects the provided data inputed by the user and stores in a single sqlite file.
With that data, you can get some insights and know yourself, then, keep learning.

##### Development stage

Know yourself and happy hacking!

---

### Entities

- **EVENT**: Every operation is marked as a global event, which is a ref to the below types

- **FACT**: Any default life event to be registered

- **COUNTER**: Used to increment or decrement like a counter. Each count corresponds to a Fact, however, the Counter itself is not a Fact.

- **HABIT**: A scheduled routine activity to be set as a target and marked as done. It inherits from Counter.

### Requirements

**E.01** - Me, as a person who wants to know more about myself, I would like to register any kind of event or data in one private, secure and resilient platform. In order to keep collecting some infos and get some insights about my personality and my life.
