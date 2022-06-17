import requests  # library to handle requests
from bs4 import BeautifulSoup  # library to parse HTML documents
import pandas as pd  # library for data analysis

import networkx as nx

import tikzplotlib

r_countries = requests.get(
    'https://simple.wikipedia.org/wiki/List_of_European_countries')
if r_countries.status_code != 200:  # means that the scrapping is not legally possible
    print("Error: website doesn't allow to download the source code")
    quit

soup_countries = BeautifulSoup(r_countries.text, 'html.parser')
table_countries = soup_countries.find('table', {'class': "wikitable"})
for data in table_countries(['sup']):
    data.decompose()
df_countries = pd.read_html(str(table_countries))
df_countries = pd.DataFrame(df_countries[0])
df_countries = df_countries.drop(["Area(km²)", "Population(As of 2011)",
                                  "Population density(per km²)"], axis=1)
df_countries = df_countries.rename(
    columns={"Name of region andterritory, with flag": "Country"})
df_countries = df_countries.sort_values(by=['Country'])

# print(df_countries)

r_borders = requests.get(
    'https://en.wikipedia.org/wiki/List_of_countries_and_territories_by_land_borders')
if r_borders.status_code != 200:
    print("Error: website doesn't allow to download the source code")
    quit

soup_borders = BeautifulSoup(r_borders.text, 'html.parser')
table_borders = soup_borders.find('table', {'class': "wikitable"})
for data in table_borders(['sup']):
    data.decompose()
df_borders = pd.read_html(str(table_borders))
df_borders = pd.DataFrame(df_borders[0])
df_borders = df_borders.drop(
    ["Total length of land borders", "No. of distinct land borders", "No. of distinct land neighbours"], axis=1)
df_borders.columns = ['Country', 'Borders']

print(df_borders)

euroGraph = nx.read_adjlist(
    "/Users/sltkaguya/Documents/itmo/discr_math/05_graph_theory/eurograph.adjlist", delimiter="--")
nx.draw_planar(euroGraph, with_labels=True)

tikzplotlib.save(
    "/Users/sltkaguya/Documents/itmo/discr_math/05_graph_theory/eurograph.tex")
