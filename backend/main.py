routes = []
indexes = {}
counter = 0

with open("routes.txt", "r", encoding="utf-8") as file:
    for line in file.readlines():
        f, to = line.strip().split("|")[:2]
        if f not in indexes:
            indexes[f] = counter
            counter = counter + 1
        if to not in indexes:
            indexes[to] = counter
            counter = counter + 1
        routes.append((indexes[f], indexes[to]))

with open("graph.txt", "w+", encoding="utf-8") as file:
    file.write(f"{len(indexes.keys())}\n")

    i = 0
    while i < counter:
        file.write(f"{i}\n")
        i = i + 1
    for route in routes:
        file.write(f"{route[0]} {route[1]}\n")