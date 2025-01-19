# Manage torrents scripts

## List removable torrents

The purpose of this script is to help to identify the torrents that are no more useful, the design goals are the following:
* Respect tracker rules (minimum seed time)
* Do not penalize torrents with few seeders or with leechers
* Do not penalize torrents with activity
* help maintain a ratio (the required ratio decreases with the number of seeders and time following a gaussian curve $f(x, s) = A \cdot e^{-((B \cdot s + C)\cdot x)^2}$ x is the number of days past the minimum time to seed and s the number of seeders)

The configuration is a json file (default `config.json` in the working directory but the path can be chosen with the `CONFIG_PATH` environment variable)

| key              | type   | default   | comment                                                                          |
|------------------|--------|-----------|----------------------------------------------------------------------------------|
| hostname         | string | localhost | the hostname of the transmission server                                          |
| username         | string |           | the username of the transmission admin                                           |
| password         | string |           | the password of the transmission admin                                           |
| port             | int    | 9091      | the port of the transmission server                                              |
| minimum_seeding  | int    | 1128      | the minimum time to seed (hours)                                                 |
| minimum_seeders  | int    | 2         | the minimum number of seeders to maintain                                        |
| minimum_leechers | int    | 1         | the minimum number of leechers that prevent to delete                            |
| last_activity    | int    | 336       | the interval (hours) to wait since the last activity before being able to delete |
| ratio_tolerance  | float  | 0.05      |                                                                                  |
| ratio_coef_a     | float  | 5         |                                                                                  |
| ratio_coef_b     | float  | 0.006     |                                                                                  |
| ratio_coef_c     | float  | 0.01      |                                                                                  |
