for v in range(1, 31):
    m = []
    if v % 3 == 0:
        m.append('fizz')
    if v % 5 == 0:
        m.append('fizz')
    print(''.join(m) or v)
