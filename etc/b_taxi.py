import sys

line = sys.stdin.readlines()


def main(line):
    numTaxi, distance = line[0].rstrip('\r\n').split()
    numTaxi = int(numTaxi)
    distance = int(distance)
    line.pop(0)
    min_charge = 0
    max_charge = 0
    for i in range(len(line)):
        l = line[i].rstrip('\r\n')
        datas = l.split()
        start = int(datas[0])
        charge = int(datas[1])
        add_distance = int(datas[2])
        add_charge = int(datas[3])

        add_result = 0
        charge_result = 0
        
        if distance >= start:
            remain_distance = int(distance - start)
            add_count = (remain_distance + add_distance) / add_distance
            add_count = int(add_count)
            add_result = int(add_count * add_charge)
        
        charge_result = int(charge + add_result)

        if min_charge == 0:
            min_charge = charge_result
        elif min_charge > charge_result:
            min_charge = charge_result
        
        if max_charge == 0:
            max_charge = charge_result
        elif max_charge < charge_result:
            max_charge = charge_result
        
    print(f"{min_charge} {max_charge}")


if __name__ == '__main__':
    main(line)
