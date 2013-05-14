#!/use/bin/env python
#coding=utf-8
from pygments import highlight
from pygments.lexers import get_lexer_by_name
from pygments.formatters import HtmlFormatter
import sys


inDir  = 'infiles/'
outDir = 'outfiles/'
def highlight_code(code, tp):
    lexer = get_lexer_by_name(tp, stripall=True)
    formatter = HtmlFormatter(linenos=False)
    return highlight(code, lexer, formatter)

def parse_file(filename, tp):
    with open(filename, 'r') as f:
        code = f.read()
        return highlight_code(code, tp)

def main():
    infile_name = sys.argv[1]
    try:
        file_name, tp = infile_name.split('.')
        outcode = parse_file(inDir+infile_name,tp)
        with open(outDir+file_name, 'w') as f:
            f.write(outcode)
    except ValueError:
        pass

if __name__=='__main__':
    main()
