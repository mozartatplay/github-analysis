<!--This File was Autogenerated from README_template.md on April 10 2018. Do not modify this file -->

README (DAVID VERSION)
=====================

This project finds trends in popularity amongst programming languages, by analyzing over 1.25 billion events from the public GitHub timeline and figuring
out how many users each language has. See [this blog post](http://www.benfrederickson.com/ranking-programming-languages-by-github-users/) for an overview and discussion of the trends.

On April 10 2018 the rankings of each language by the number of active users on GitHub are:


<table>
<thead><tr><th>Rank</th><th>Language</th><th align="center">MAU</th><th align="center">Trend</th></tr></thead><tbody>
<tr><td>1</td><td>JavaScript</td><td align="right">21.41%</td><td align="center"><img src="./images/JavaScript_sparkline.svg"></td></tr>
<tr><td>2</td><td>Python</td><td align="right">14.69%</td><td align="center"><img src="./images/Python_sparkline.svg"></td></tr>
<tr><td>3</td><td>Java</td><td align="right">12.65%</td><td align="center"><img src="./images/Java_sparkline.svg"></td></tr>
<tr><td>4</td><td>C++</td><td align="right">8.30%</td><td align="center"><img src="./images/C%2B%2B_sparkline.svg"></td></tr>
<tr><td>5</td><td>C</td><td align="right">5.84%</td><td align="center"><img src="./images/C_sparkline.svg"></td></tr>
<tr><td>6</td><td>PHP</td><td align="right">5.31%</td><td align="center"><img src="./images/PHP_sparkline.svg"></td></tr>
<tr><td>7</td><td>C#</td><td align="right">4.79%</td><td align="center"><img src="./images/C%23_sparkline.svg"></td></tr>
<tr><td>8</td><td>Shell</td><td align="right">4.78%</td><td align="center"><img src="./images/Shell_sparkline.svg"></td></tr>
<tr><td>9</td><td>Go</td><td align="right">4.36%</td><td align="center"><img src="./images/Go_sparkline.svg"></td></tr>
<tr><td>10</td><td>TypeScript</td><td align="right">3.82%</td><td align="center"><img src="./images/TypeScript_sparkline.svg"></td></tr>
<tr><td>11</td><td>Ruby</td><td align="right">2.94%</td><td align="center"><img src="./images/Ruby_sparkline.svg"></td></tr>
<tr><td>12</td><td>Jupyter Notebook</td><td align="right">2.66%</td><td align="center"><img src="./images/Jupyter%20Notebook_sparkline.svg"></td></tr>
<tr><td>13</td><td>Objective-C</td><td align="right">1.77%</td><td align="center"><img src="./images/Objective-C_sparkline.svg"></td></tr>
<tr><td>14</td><td>Swift</td><td align="right">1.67%</td><td align="center"><img src="./images/Swift_sparkline.svg"></td></tr>
<tr><td>15</td><td>Kotlin</td><td align="right">1.11%</td><td align="center"><img src="./images/Kotlin_sparkline.svg"></td></tr>
<tr><td>16</td><td>Rust</td><td align="right">0.86%</td><td align="center"><img src="./images/Rust_sparkline.svg"></td></tr>
<tr><td>17</td><td>R</td><td align="right">0.78%</td><td align="center"><img src="./images/R_sparkline.svg"></td></tr>
<tr><td>18</td><td>Scala</td><td align="right">0.74%</td><td align="center"><img src="./images/Scala_sparkline.svg"></td></tr>
<tr><td>19</td><td>Lua</td><td align="right">0.68%</td><td align="center"><img src="./images/Lua_sparkline.svg"></td></tr>
<tr><td>20</td><td>PowerShell</td><td align="right">0.53%</td><td align="center"><img src="./images/PowerShell_sparkline.svg"></td></tr>
<tr><td>21</td><td>Matlab</td><td align="right">0.47%</td><td align="center"><img src="./images/Matlab_sparkline.svg"></td></tr>
<tr><td>22</td><td>CoffeeScript</td><td align="right">0.44%</td><td align="center"><img src="./images/CoffeeScript_sparkline.svg"></td></tr>
<tr><td>23</td><td>Perl</td><td align="right">0.43%</td><td align="center"><img src="./images/Perl_sparkline.svg"></td></tr>
<tr><td>24</td><td>Groovy</td><td align="right">0.37%</td><td align="center"><img src="./images/Groovy_sparkline.svg"></td></tr>
<tr><td>25</td><td>Haskell</td><td align="right">0.37%</td><td align="center"><img src="./images/Haskell_sparkline.svg"></td></tr>
</tbody></table>


### Data Sources

The main data sources for this project are:

 * The [GitHub Archive](https://www.githubarchive.org/) project which has been recording every public event on GitHub since early 2011. Overall there are more than 1.25 Billion events stored there, which are more than
 400GB gzipped.
 * The [GHTorrent](http://ghtorrent.org/) project. Which also monitors the GitHub public event timeline, and retrieves extra information from the GitHub api for each event seen.
 * A custom scraper I wrote here, which backfilled missing repo information from the GitHub API.

### Inferring Languages

Analyzing programming language trends requires figuring out the language for each repository.

There are two ways to get language information out of the GitHub API. The first is to query for the Repository using the [GET /repos/:owner/:rep](https://developer.github.com/v3/repos/#get) API.
This returns the dominant language of the repo. You can also get the byte breakdown of all the languages using the [GET /repos/:owner/:repo/languages](https://developer.github.com/v3/repos/#list-languages) endpoint. This will include the number of bytes used by each language in the Repo.

We're using the single dominant language for the repo for this analysis. While this loses some information, there are multiple benefits that make this much more practical:

 * All GitHub Archive events from 2012/03/10/ to 2014/12/31 included the repo language in the repository.language json field.
 * All PullRequestEvent events have this information in the payload.pull_request.base.repo json field
 * The GHTorrent project has 40 million repos with languages in the projects table, but only has the detailed language breakdown for 27.6 million repos in the project_languages table

So the plan to get the language for each repo is to aggregate all these sources of information: The language scraped from the GitHub REST API,  the language from the projects table
of the GHTorrent project, the language included with certain Github Archive events and finally the language inferred from fork events (forks are assumed to have the same
language as the repo they are forking)

| Source         |  Repo Count   |
| -------------  |:-------------:|
| GHTorrent      | 42.8M         |
| scraper        | 22.5M         |
| GitHub-Archive | 13.5M         |
| forks          | 37M           |

There is significant overlap between all these sources of information, but once aggregated and
deduplicated there we ended up with language information for 62.8 Million repos. This includes every repo that has had more than 1 user interact with it and all repos with only 1 user that have more than 5 total events ever.
I'm still crawling the remaining repos.

### Inferring the Date

The dates here are the dates corresponding to the events in GitHub Archive. This means that we are analyzing when something was pushed to GitHub rather than the commit date (which could be considerably earlier).
The reason behind this is that the commit date can potentially be inaccurate: The dates are given from the developers and it's not uncommon to see dates that implausibly far back in the past (Jan 1, 1970), or even more
 implausibly occurring in the future.

### Running the Code

This code requires both Go and Python to run properly.  Additionally, this requires around 1TB of free disk space to run, and I would recommend at least 16GB of RAM.

To configure your system to run this code
  * Install all the python dependencies by running ```pip install -r requirements.txt``` and go dependences by running ```go get ./...``` from this directory
  * Install Postgres onto your system, and create the database tables from the schema.sql file: ```psql github < schema.sql```
  * Copy the config_template.toml file to config.toml and fill out the required fields.

There are multiple different components to this code.

The main programs written in Go are:

 * ```gha-download-files```: downloads new files from the githubarchive so that they can be analyzed locally.
 * ```gha-parse-events```: Parses the JSON events from the Github Archive and converting to normalized TSV files.  The JSON event schema changes several times over the last 7 years, and normalizing to a consistent TSV schema makes it much easier to analyze.
 * ```gha-scraper```: Crawls repo information from the GitHub API and inserts into Postgres.

 There are also several small bash scripts that do the actual analysis:

 * ```scripts/calculate_language_mau.sh```: Joins the repo languages against the parsed events, and figures out the MAU for each language at every month.
 * ```scripts/calculate_repo_languages.sh```: Merges information from postgres/ghtorrents/extracted GitHub archive events/ and from fork events to get a single repo:language mapping.
 * ```scripts/calculate_top_repos.sh```: Ranks each repository by the number of users. The output of this is passed to gha-scraper to crawl repositories.

 Finally plotting is done with Python by running ```python scripts/plot.py```. This will also update the graphs in this README.

A future goal of this project is to simplify the steps needed to run this code, it's unnecessarily convoluted right now.
