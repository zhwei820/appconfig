from jinja2 import Template
import os


def read_file(src):
    struct_start = 0
    struct_end = 0
    kk = 0

    lines = []
    title = ''

    with open(src, 'r') as f:
        s = "nullstr"
        while s:
            s = f.readline()
            if struct_start and 'orm' not in s and 'json' not in s:
                struct_end = 1

            if 'struct' in s and 'type' in s:
                struct_start = 1

            if struct_start and not struct_end:
                if kk:
                    lines.append(s)
                else:
                    title = s
                kk += 1

    return lines, title


def get_db_field(line):
    pos1 = line.find('column(')
    pos2 = line.find(');', pos1)
    db_field = line[pos1 + 7:pos2]
    return db_field


def parse_model(src):
    lines, title = read_file(src)
    title = title.split(' ')[1]

    items = []
    for line in lines:
        bean = [item.strip() for item in line.split(' ') if item]
        bean = bean[0:2] + [''.join(bean[2:])]
        items.append(bean)

    model_data = {
        'obj': title.lower(),
        'objBig': title,
        'data': [],
        'testInput': "",
    }
    for item in items:
        db_field = get_db_field(item[2])
        field_type = 'int' if 'int' in item[1] else item[1]
        if any([x in db_field for x in ['id', 'create', 'update']]):  # 这些字段数据库生成
            continue
        model_data['data'].append({
            'db_field': db_field,
            'model_field': item[0],
            'field_type': field_type,
        })
        model_data['testInput'] += '''"%s": %s,''' % (db_field, '"%v"' if field_type == 'string' else '%v')
    model_data['testInput'] = model_data['testInput'].strip(',')
    return model_data


def render(model_data, tpl=''):
    with open(tpl, 'r') as f:
        s = f.read()

    template = Template(s)
    t = template.render(model_data=model_data)
    return t


def write_result(t, out):
    os.makedirs('/'.join(out.split('/')[0:-1]), exist_ok=True)
    if os.path.isfile(out):
        raise Exception('%s 已经存在,无法自动生成' % out)

    with open(out, 'w') as f:
        f.write(t)


def work(src):
    model_data = parse_model(src)
    t = render(model_data, tpl='utils/template_service/crud_api_test.go.tpl')
    out = 'services/%s_service/crud_api_test.go' % (model_data['obj'])
    write_result(t, out)
    print('生成: %s' % out)

    t = render(model_data, tpl='utils/template_service/crud_api.go.tpl')
    out = 'services/%s_service/crud_api.go' % (model_data['obj'])
    write_result(t, out)
    print('生成: %s' % out)


if __name__ == '__main__':
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument("src", nargs='?', default="group.go")

    args = parser.parse_args()

    try:
        work(args.src)
    except Exception as e:
        print(e)
        exit()
    print('完成!')
