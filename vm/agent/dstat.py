import string

with open('dstat.dump', 'r') as f:
    lines = f.readlines()
    old_indexs = lines[1].split()
    new_indexs = [string.strip("-") for string in old_indexs]
    sub_indexs = lines[2].split("|")
    sub_indexs = [string.strip() for string in sub_indexs]
    index_dict = {}
    for index in range(len(new_indexs)):
        sub_index = sub_indexs[index].split()
        sub_index_dict={}
        for i in range(len(sub_index)):
            value_dict=[]
            for j in range(3,8):
                values = lines[j].split("|")
                try:
                    value_dict.append(values[index].split()[i])
                except:
                    value_dict.append(None)
            sub_index_dict[sub_index[i]] = value_dict
        index_dict[new_indexs[index]] = sub_index_dict

    print(index_dict)

