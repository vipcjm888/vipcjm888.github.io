# Process the flybook message card from JSON to a line of text


def process_text(input_text):
    processed_text = ""
    inside_quotes = False
    
    for char in input_text:
        if char == '"':
            inside_quotes = not inside_quotes
            processed_text += '\\"'
        elif char == ' ' and not inside_quotes:
            continue
        elif char == '\n':
            continue
        else:
            processed_text += char
            
    return processed_text

# 输入文本
input_text = """

"""

# 处理文本
output_text = process_text(input_text)

# 打印处理后的文本
print(output_text)

