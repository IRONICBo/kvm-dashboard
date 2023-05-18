with open('perf.dump', 'r') as f:
    lines = f.readlines()
    index_dict = {}
    for i in range(1, len(lines) - 2):
        time_dict = {}
        line = lines[i].strip().split("|")
        if line[2] == "":
            new_dict = dict(pid=line[1], count=line[4], time_enabled=line[5], time_running=line[6], metric=line[7])
        else:
            new_dict = dict(count_num=line[1], unit=line[2], count=line[4], time_enabled=line[5], time_running=line[6], metric=line[7])
        time_dict[line[0]] = new_dict
        try:
            index_dict[line[3]][line[0]] = new_dict
        except:
            index_dict[line[3]] = {}
            index_dict[line[3]][line[0]] = new_dict
    print(index_dict)
