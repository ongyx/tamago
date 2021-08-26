regs = {
    "B": "BC.Hi",
    "C": "BC.Lo",
    "D": "DE.Hi",
    "E": "DE.Lo",
    "H": "HL.Hi",
    "L": "HL.Lo",
    "(HL)": "",
    "A": "AF.Hi",
}

instructions = ["bit", "res", "set"]

struct_tem = """
// {opcode}
{{
asm:	"{asm}",
length: {length},
cycles: {cycles},

fn: func(s *State, v Value) {{
{body}
}},
}},
"""

body_tem = """
s.fl.{ins}({pos}, &s.{reg})
""".strip(
    "\n"
)

body_hl_tem = """
b := s.ReadFrom(s.HL)
s.fl.{ins}({pos}, &b)
s.WriteTo(s.HL, b)
""".strip(
    "\n"
)

structs = []

opcode = 0x40

for ins in instructions:
    for bit in range(8):
        for name, reg in regs.items():
            if name == "(HL)":
                body = body_hl_tem.format(ins=ins, pos=bit)
                cycles = 3 if ins == "bit" else 4
            else:
                body = body_tem.format(ins=ins, pos=bit, reg=reg)
                cycles = 2

            struct = struct_tem.format(
                opcode=hex(opcode),
                asm=f"{ins.upper()} {bit},{name}",
                length=2,
                cycles=cycles,
                body=body,
            )
            structs.append(struct)

            opcode += 1

open("cbops.table", "w").write("\n".join(structs))
