package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/4SYdzAcatvGVy8JQbSwrAii97CEKEFEO5ZFdDZEDdEiIlmsE2Gj915WCaIFq4RdLbqsVUL0qKsTrF6+Oe8773fO+ead891/PHM/M/dccz3PP9dvbkNdMnJmADWAGgB7BjUC/KVYATSAp96O9r6O4j5e7uISf/YmUAoASQOFlN2m5Y7nSBdnkG3JVQfGnV84acdfO8PQ9vclC14jKqBZqNlhJ1aAspue4KM7F0IZ+ZxDEzEkBafG88UKIAXYDRF1QQIGk1cwxSvt4a8/Srjuik3u077BEk6mfXIGlf0JteQUrY7J9jRPfIkbFiefEkaMR4w312Bx70/4DZTIJMInRiQDi8hZW5tHlOSpVtQHAtRUHenIEph0i1jb7L7zBLNzcKP8oyWyKdRXoCSIc95gwj5/HjnfmLmignLSsr1WsE3TGXNw2prQrnkCVJPUEY1aZy57UXjOROUTLEIbbCnErPpMUxAZIN7f4dhD063O1Ut+B8nCCKERDIeVju0OXDlZP1JoSDUXarDpmHni58jswVw2z9gwBKHTVZU5oO00E+zJxu2mlQECWp/8/Ezq8OEdqW7TUZDXZcZJAGMJqkzrEbifkCJzTvGC0gYg4y79rE1HaFAsxJj0hr3zuFJKKGGC3IZxQu+D8sPsNa6rBzEjxjUAb90AegRF7LjfHf9Sz+Q2UvCsPP5dSJ44OV2yBFxEnPVk+fpG3BQeop9006VdWorxJk3+or0h10wRdQTgffLmlmT57wZJi+1pBRrihwmd733LNl6/2gsWDzGViMiyaqBJi1iTtFzSIc3p9VGPNQv6fvDT/+HuC8a92H71IP79JMMlVXrWXADJDV3SnEHazaG+CasiAuQNdNP2ovi8YPvRMjaAAaqmeF8JxYE6IhCZRQsFWS7MnNuHUlMLbNwLFKW/+YUKrmmpuYXz8obc01gS5pXpJhAITosUHA9C84UKos9bmtVdjUKLfwDb1ao2MZehfUecvLrFd4KKLsDxVKqwaiZ9pCsRJ+wtCfeV83kgcuW/oLFeQO4RJqdBkSEN6X96A1SBvNNpr1Es7EajvXmf3cHrI3981+OAgXA99OeRXBmNId7JH6FBanxQt+L4YITHkRrp68H81ZTINx7McR6jnwnLdjfXiCCVUHpcTodaMkdFklTEIGvlFY/1eZgglo9v4HfhIk3Zz0HXFIrwRVb3rutPt4q+vB5RtdjqQIvpjjhep2Tq1K6oHqnEukZyXkKSwPcN+UupsaFqcf4WvCbZ0JWsJ4vZRPYRuLM/pR9MMFnvpmSbD0MreYvjdb3Ur8gbfPqFo41t2QMLCo2OGY9jDLliMKMKZV/CHdNzYXrV0KbPmKWIdMitwF5vxsS4NyXq13XNBG7OJdJ0fTG/t+qwMB5xuy9MGj4PbPYFRivmrTMn3GOp7K5ktmYjudvrYemL1raR6PoVMUq3OHv6Aeo4BJsFwGlqH3vnbvoxqWflRSdH838cLKgDGizhwjtxYNiJlJR5x3XsNjj9q+GYV6Mre4rAi/uLJIaSIl33yu5OQsAVeshiR4HdF1BNgyqX70xAAcwlSQv5HgeqBU/JmRFxzOrktToRd2R5btQxMI+Hv/fxycW8/Y59O4IN6JiddZm5moWjWWg9TIiVw2fPR4ef/jzktvKUmzXXuNbJx5keJvKkLO8DZZxR98HxyfaG7cvK2R+Bu80XRUPnL+pP5usN32xLgxxKFmpsp9E750BkGcKPIoHH3bEDLzlMZhEmRZWGMIsrhORapteBzDfWr52NgeZ7qI0QCcmGcYwKUwFTl8e86tlv+992G6kdzvRjylO2VDgPE3LU7edF/PSgZf1k5rr+wqAL/OTrL0fuHy9eNM38qPWcu675+HPYhsnDbpTH9z7MqK4J7VZg0BYW7u6J9VTBsCJ39pZJ7LILkTvTCoH7bwWr/ETNGOlIU4Qlz3a2VlQGX1bOuLW41G7v/MiR69hYxT53rCxE0VEV9iub3KJPPbLsqiikGpi9vEeaFOHXeX+cmOY/8MGlrTgdBHeHv14fOxsba8KU1zMKGZcgzwNZTSrD3SUKIUJJiRP13PW2yu7y23XZSoo8oioh+ThzUOnYStZK99x7nUwJUhV7LH8wWxs6j4nQ0e6mXcNgQjfzLnmBdu/0dQ0jtsrPnnjNfnR/4h2LcsQjItCkg6al74Qe8QBz48xEfH4mnI9TsVNR/8RVn5ddCOW/EO3zk7u3c7ogwvR0eu4AqHLiCQAArq4MdamoCwG3ECyUAACFLADwn4QBANa1TP6WMLR/S5h/pcoJSNruD4G/jv0ZUVH/R4AJQAPwdfTx/XdA/af7l5AuSNpuIM3WmMP8tHvHOOkIz2asAr/f6ZozcuXvJmFUM2ZQVz0x8dXh2dEzvrgoPWYg/rMAm2VqOG8uIn9IMHoh3I7XjikKH4xlGyH4HiV52m/4hqo/jfeQmUj/eu3dytkblcERQibBpWKuecbnU8oXeM3RDtTWmY/hIODh29EZN/ntg+miS1CDUCuJqg+wwFx+uX2o4JR329tYHPBC8bWFrvPEU795gKcD3pVEX2s0PfLQCKlUGSVRTEzEKkSMjbOYd43IOGMQvCZoxTzTjwZuNLLK+b0Sr68Q9DU6ylFIS4s1Jz/7GWKTig1mD24cmwpT4JEOJvuKBOTa0yz4MmuOJ3rrJcUQC5WiuUuJiacLt2jkaxDIDKeg9TTaaFGUg4GYXsYKUU/TWixehC4HqVSyvS28LGlQmd8upyM29aKLfW3hVs8vpt+AEbKSmfO+UCG23eFtl+cU3+7uzusmwqNu0BPNHGvFN4yFCusQ8ZDCbAFKV3vZFPVAn9Lmkk6beNs9L77BNPq5MFC+tulcCaocOFsp7E3q5PYRTR7t5PXaImqnVGJB8FAGuYlaGJvafdpx+F3EH3u6j57lLvWUXaI7hNQHvi7B2v8mKMtKfXgECUJZ5agNOgHgK0fEMzXtJTLZhqNO+Z+HoU2XQ8H+ZWefgma3TzfISuRX2S79tkA3MP6e1i3joNO9bnGG8+1HbERiinI9xwJ8MHbc8ODgxsm5+J1XlmNKbYxzF2r91itfAj++Hx6XQw+jsI163nKS0tLA04f+R/khkJgg/OLOa9mfnOxbX0Ds69zblGU9NQyqdTWCWXt4/UYPylMMs2K089QiONPtO4iT0y3uEj7KyfGL4nzGEHY79kn0xuHxw6ezdvIiEaLMztGefLVixZ49+63q5JawNzv7Z1pNXW9iuN9paZ29JbkzGuu/mi1/HNhKpq0b/5VWdaEJywnHRziYvMeNkuiP8HpaZnkV44aCw3jdjLP8jFvjFCW8EXyTUMI085a2NcycVR2XrbhjaB0brLofXPdRHDxwHZIC8K/wGJKvRIvKJGaa2CYaSofQTrU6ygPvqlv85I9e5BujJhTx0Gl483xtB1sP5yzp26dkwKYTP92s67hemxpi8E2VJry98VabUlbPg+dTZE9/tYKJPNKi13DK2h59hdCnzm5UXJpeUhWP6H/tiU5Otn06W8kRVHDDfuL+hnKfwZp4H6PW0nTHo6nKeoIPJ4WWq4zgHqWlTQF+8jvSnmqccPAH0N1lMsRXVUG19B2Ow/rEmGpQdv3v/kYKTNl5t3okU+h/EzdbAp66TLvcezWs80tWTBvlXBJWUn/EJJFSu+VV3pAvXrZ4IWwzI+/JWst+v/yqTB9jumRe/qqARcYCjEjR1blXBkW9WmRo6VuNCYIIQ966ak1xP+oFKlxGzZyvZHHKP4jkt6jaZPveusDHodrq+1Z+87k5wysfTpWAqPzHhaJuVDfBmA9Iz8yGeNGKzfFJX/49UjK1DyKIkgrH3dHm5Ay5DKXETC8YTexT1fsGrr2byx6A/aAzXAhsEOJuYp3Gk+UlkPfujrh852QduWjyJtugFROwfz9W2ZZpIS37ki+ormbyiEwMvplS/lmDn5FWTBfqKgqZomev1kFjEKzcj5aLQ1cdWhxcDBDJNRVS62VbeUJpA9d9HBh4/RFfsjV7b9Wi80djvNwh91MNHRotb1QKdt/yTeWfCbS3DKqLQH06XuNGVggvs0BLGc3kJCOfvnx+6M1+/c6v5Kgiv+PWD4O13ZW4R7Grwn2dLOZwTm6zhATkNN9odXcxuFwcmKtu9K3whWhknbeDqKv9WhEHmb6I46GQWo6O0LKwC8YpfrKN8jE9civ3R6lRcpKO1VzmKGoQDaPUJPTyyKRAo02HTA1y0o0/+Zzukxvo5Rl9M6iIpLt7sWfUCCPXSjw+tEJ/Tr2XKb42KEfsXNR6pliJzyFAlpPWJmTVn30pZ2nOkLKVJYB9ppuKve74pA1x2GCGJ78Q3/hsedrOrOale9oeD9nlBh1chhPzbV9TUZHZUiFCL0bFYq5uP460CCxdWbQoyV3uePXSpp9v92V6c8iKjmCWSslW6kNCL8Q07i56WqL4MsEoetE31SiH2aD0mmfHIphrjN704CRlve9r2mzvEOktSAJN80ERpK8nss729niz4nDHyF30j8d+aiJcJSfgXr1GbY+w9jqjiarwhoFCA7mVT2bFXS1HKbVxwc/0b9+Rym0sSbPYuTpXHrx4HyW3Vn3bXMshErfi4wRavP0bDkeVUx+GffGpKbcPv0moDX3YnKyJrrLRbxxIOXiHEvZc9L71iBB5ynHNzPrJo99GTo1MSvjQ0w1fpPO4Vn5pacXV462AJjb5jZgRi/ZzJR3oQU+knIHilIzY/sxYcLEVdJ1zxfHKyaYgsT5xL9JU+IKw6NaLly2geI924NiYaMDKlI/FDags0F3+3I82Z2MouK1PkhmdHiOo/bGHhomMPI5wn6rpA9TOWNPZLA6c/q7ZmqULSsLsbjxP0kEW5/TzLLXHYMLA8SW+6LqweGgXmGmPMmiXcufGq/O509NeMZkVjDwDtdL5S1mw82HPzxmPqfxWnODueIz3g24FxxB6VFUWxL0aHNGe5ddkYM5OBTnGCcqZZ2ow4iy134uWhQrCnJgzRA1QsbgFnN9HYGYodZfhEFO7gIbM7/zaeE4OLted287VK99h7N8MUL3UwqYPlNC1wPeJuiLPRq10guoew4XE52v1bIiBYc0aq3PZbbH1EyVk8veYgESOuZt48cUJpZmXKNIe/JvFl64zrRYVeTNRIDYXtGL1K0aeWgD8FSVvcBWZYdasVEuLMynT7Zur40s11ZNUjqK5IO4KixsP2RSF400Pi51fHujrrIwk/tpLLaNIQL+bAE5ZrbeTCBW1OePs9Y0OE0ZHpEBO84vuv7z0sTiaGiJL4XmCkY0Lk8/OXY+Hzx1y3x2fb7Ee9DPss8JmWaWuPGDL4NBuXTbG/RY6WJu66A9e1ln9su2e02Fsb3mi068FbmHRtMH5qZs8b+wUriqleSykeSb7gl/VITeElsE+NusLssiljMiF7ufmbNLw0fhk0ap4sSclZKblH0ddv3CZnupvWmjlDI9p4evYOedCrQETdfc4XXMO6p6H1HdODclNbGaIO2eu8xeedrVfe1BaGu11PzvP+rhq7UNw0NcNkBTIG41LuaL7D/SZcw2Ru9EAAJK8/wR91H9htn9APhJSZrI/9xJ/5UlWAM3/ztWF/HH+P1uKP8X+O4L+uxgAV2owSsB/B9I/3fyBoH99KNPf3ORTAv4JSP/JCfXfnJgxAP7bL1FQ/jFADiAHsJIAAMkMf9z+JwAA//+h1CVDwhEAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
