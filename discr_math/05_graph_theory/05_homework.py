import json
import networkx as nx
import matplotlib.pyplot as plt
from networkx.readwrite import json_graph as nxjson

euroGraph2 = nx.read_edgelist("05_edge_list2.txt")
euroGraph2.add_nodes_from(["Cyprus", "Iceland", "Malta"])

euroGraph = nx.read_adjlist("05_edge_list.txt", delimiter="--")
nx.draw_spring(euroGraph)
