# #mobile-squad (Slack) — Thu 10 Sep

**Mobile dev** 14:02
Heads up — finance asked us to stop sending money as floats because of rounding
(0.1 + 0.2 = 0.30000000000000004 etc). For v2 we're sending all amounts as decimal
strings: `"amount": "45.00"`. Backend should be fine, it's still JSON 🤷

**Mobile QE** 14:05
👍 I'll update the app-side tests

**Mobile lead** 14:11
Nice. Does anyone need to tell the claims squad?

**Mobile dev** 14:12
it's still the same fields, just a string instead of a number. should be non-breaking

*(no further replies — the claims squad are not members of #mobile-squad)*
