from TexSoup import TexSoup, utils
import json
import re

invalid_tag_parse = ["equation", "bibliography", "bibliographystyle"]
special_node = ["figure", "table"]


def parse_node(node, pure_english: list):
    # base case, process the node
    if isinstance(node, utils.Token):
        if not str(node).strip().startswith("%"):
            pure_english.append(node)
    else:
        if node is not None and node.name in special_node:
            parse_special_node(node, pure_english)
        elif node is not None and node.name not in invalid_tag_parse:
            for temp_node in node.contents:
                parse_node(temp_node, pure_english)


def parse_special_node(node, pure_english: list):
    """
    handle table, figure node

    the two above node need to parse only one node named caption

    :param node: table or figure node
    :param pure_english: total parsed tokens
    :return: None
    """
    if node.name == "figure" or node.name == "table":
        for temp_node in node.contents:
            if not isinstance(temp_node, utils.Token) and temp_node.name == "caption":
                pure_english.extend(temp_node.contents)


def extract(content):
    soup = TexSoup(content)
    all_contents = []
    pure_english = []
    # extract the title part
    for title_part in soup.contents:
        pass

    parse_node(soup.document, pure_english)

    return pure_english


def get_line_number(components, content):
    all_lines = content

    # list of result [[string, line_number]]
    result = []

    # search index in the file
    line_number = 0

    for one_component in components:
        search_item = []
        # check the sentence is multiple lines
        if '\n' in one_component:
            temp_search_item = one_component.split('\n')
            for temp_item in temp_search_item:
                if temp_item:
                    search_item.append(temp_item)
        else:
            search_item.append(one_component)

        # search all the components using the line number
        for item in search_item:
            while True:
                if str(item) in all_lines[line_number]:
                    result.append([item, line_number])
                    break
                else:
                    line_number += 1
    return result


def process_final_result(original_result, content):
    all_lines = content
    final_result = []
    if original_result:
        # initial line number
        previous_line = original_result[0]
        flag = False
        for line in original_result[1:]:
            if line[1] == previous_line[1]:
                flag = True
            else:
                if flag:
                    final_result.append([all_lines[previous_line[1]].replace('\n', ''), previous_line[1]])
                    flag = False
                else:
                    final_result.append(previous_line)
                previous_line = line
        final_result.append(previous_line)

    return final_result


def filter_items(original_result):
    """
    remove and process some items
    :param original_result:
    :return:
    """
    # remove_index = []
    #
    # for index in range(len(original_result)):
    #     for i in original_result[index][0]:
    #         if i == ' ':
    #             break
    #     else:
    #         remove_index.append(index)
    #
    # for index in remove_index:
    #     original_result.pop(index)

    # extract the tag content
    for index in range(len(original_result)):
        temp = original_result[index]
        temp_result = re.search("^\\\\[a-zA-Z]*\{(.*)\}", temp[0])

        if temp_result is not None:
            temp[0] = temp_result.group(1)

    return original_result


def make_json(original_result):
    result = {"data": []}

    for item in original_result:
        result["data"].append({"text": item[0], "lineNumber": item[1]})
    return json.dumps(result, ensure_ascii=False)


def main(content):
    lines = content.split('\n')
    ext = extract(content)
    result = get_line_number(ext, lines)
    final_result = process_final_result(result, lines)
    final_result = filter_items(final_result)

    json_result = make_json(final_result)
    return json_result
