from typing import List, Tuple, Optional, Union


JSON_WHITESPACE = [" "]
JSON_SYNTAX = ["(", ":", ")", ",", "{", "}", "[", "]"]
JSON_QUOTE = '"'
JSON_DIGITS = [str(d) for d in range(0, 10)] + ["-", "e", "."]


def from_string(s: str) -> dict:
    pass

# lexer
def lex(s: str) -> List[str]:
    tokens = []
    while len(s) > 0:
        json_string, s = lex_string(s)
        if json_string is not None:
            tokens.append(json_string)
            continue

        json_number, s = lex_number(s)
        if json_number is not None:
            tokens.append(json_number)
            continue

        json_bool, s = lex_bool(s)
        if json_bool is not None:
            tokens.append(json_bool)
            continue

        json_null, s = lex_null(s)
        if json_null is not None:
            tokens.append(None)
            continue

        if s[0] in JSON_WHITESPACE:
            s = s[1:]
        elif s[0] in JSON_SYNTAX:
            tokens.append(s[0])
            s = s[1:]
        else:
            raise Exception(f"Unexpected char: {s[0]}")
    
    return tokens

def lex_string(s: str) -> Tuple[Optional[str], str]:
    json_string = ""
    if s[0] == JSON_QUOTE:
        s = s[1:]
    else:
        return None, s

    for c in s:
        if c == JSON_QUOTE:
            # this is the ending quote
            return json_string, s[len(json_string)+1:]
        else:
            json_string += c
    
    # if no end quote, raise
    raise Exception("Expected EOS quote!")

def lex_number(s: str) -> Tuple[Optional[Union[int, float]], str]:
    json_number = ""
    for c in s:
        if c in JSON_DIGITS:
            json_number += c
        else:
            break

    rest = s[len(json_number):]

    if len(json_number) == 0:
        return None, s
    if "." in json_number:
        return float(json_number), rest
    return int(json_number), rest

def lex_bool(s: str) -> Tuple[Optional[bool], str]:
    if len(s) >= 4 and s[:4] == "true":
        return True, s[4:]
    elif len(s) >= 5 and s[:5] == "false":
        return False, s[5:]
    return None, s

def lex_null(s: str) -> Tuple[Optional[bool], str]:
    if len(s) >= 4 and s[:4] == "null":
        return True, s[4:]
    return None, s

# parser






if __name__ == "__main__":
    s = '{"foo": [1, 2, {"bar": 2}]}'
    print(lex(s))