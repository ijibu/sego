/*字典文件中为如下数据：
中国 2 n
中国人民 4 n
中国人民共产当 7 n
*/
dict = &sego.Dictionary{
	root: sego.node{
		word:  sego.Text{},
		token: nil,
		children: []*sego.node{
			&sego.node{
				word:  sego.Text{228, 184, 173}, //中
				token: nil,
				children: []*sego.node{
					&sego.node{
						word: sego.Text{229, 155, 189}, //国
						token: &sego.Token{
							text: []sego.Text{
								sego.Text{228, 184, 173},
								sego.Text{229, 155, 189},
							},
							frequency: 2,
							distance:  2.700439691543579,
							pos:       "n",
							segments:  []*sego.Segment{},
						},
						children: []*sego.node{
							&sego.node{
								word:  sego.Text{228, 186, 186}, //人
								token: nil,
								children: []*sego.node{
									&sego.node{
										word: sego.Text{230, 176, 145}, //民
										token: &sego.Token{
											text: []sego.Text{
												sego.Text{228, 184, 173},
												sego.Text{229, 155, 189},
												sego.Text{228, 186, 186},
												sego.Text{230, 176, 145},
											},
											frequency: 4,
											distance:  1.700439691543579,
											pos:       "n",
											segments: []*sego.Segment{
												&sego.Segment{
													start: 0,
													end:   6,
													token: 0x1184A690,
												},
											},
										},
										children: []*sego.node{
											&sego.node{
												word:  sego.Text{229, 133, 177}, //共
												token: nil,
												children: []*sego.node{
													&sego.node{
														word:  sego.Text{228, 186, 167}, //产
														token: nil,
														children: []*sego.node{
															&sego.node{
																word: sego.Text{229, 189, 147}, //当
																token: &sego.Token{
																	text: []sego.Text{
																		sego.Text{228, 184, 173},
																		sego.Text{229, 155, 189},
																		sego.Text{228, 186, 186},
																		sego.Text{230, 176, 145},
																		sego.Text{229, 133, 177},
																		sego.Text{228, 186, 167},
																		sego.Text{229, 189, 147},
																	},
																	frequency: 7,
																	distance:  0.8930847644805908,
																	pos:       "n",
																	segments: []*sego.Segment{
																		&sego.Segment{
																			start: 0,
																			end:   12,
																			token: 0x1184A6C0,
																		},
																	},
																},
																children: []*sego.node{},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
	maxTokenLength: 7,
	numTokens:      3,
	tokens: []*sego.Token{
		0x1184A690,
		0x1184A6C0,
		0x1184A6F0,
	},
	totalFrequency: 13,
}

//传入的带分词句子为:中国人民共产当
//输出:中国人民共产当/n
segments = []sego.Segment{
	sego.Segment{
		start: 0,
		end:   21,
		token: &sego.Token{
			text: []sego.Text{
				sego.Text{228, 184, 173},
				sego.Text{229, 155, 189},
				sego.Text{228, 186, 186},
				sego.Text{230, 176, 145},
				sego.Text{229, 133, 177},
				sego.Text{228, 186, 167},
				sego.Text{229, 189, 147},
			},
			frequency: 7,
			distance:  1.6520767211914062,
			pos:       "n",
			segments: []*sego.Segment{
				&sego.Segment{
					start: 0,
					end:   12,
					token: &sego.Token{
						text: []sego.Text{
							sego.Text{228, 184, 173},
							sego.Text{229, 155, 189},
							sego.Text{228, 186, 186},
							sego.Text{230, 176, 145},
						},
						frequency: 4,
						distance:  2.4594316482543945,
						pos:       "n",
						segments: []*sego.Segment{
							&sego.Segment{
								start: 0,
								end:   6,
								token: &sego.Token{
									text: []sego.Text{
										sego.Text{228, 184, 173},
										sego.Text{229, 155, 189},
									},
									frequency: 2,
									distance:  3.4594316482543945,
									pos:       "n",
									segments:  []*sego.Segment{},
								},
							},
						},
					},
				},
			},
		},
	},
}
