import networkx as nx
import matplotlib.pyplot as plt

euroGraph = nx.read_adjlist(
    "/Users/sltkaguya/Documents/itmo/discr_math/05_graph_theory/eurograph.adjlist.txt", delimiter="--")
nx.draw_planar(euroGraph, with_labels=True)
plt.show()
