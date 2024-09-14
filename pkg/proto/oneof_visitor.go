/*
 * Copyright 2023 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package proto

import (
	"strings"
)

type OneOfVisitor struct {
	Visitors []Visitor
}

// NewOneOfVisitor creates an OneOfVisitor
func NewOneOfVisitor() *OneOfVisitor {
	Log.Debug("Initializing OneOfVisitor")
	out := &OneOfVisitor{Visitors: make([]Visitor, 0)}
	out.Visitors = append(out.Visitors,
		&CommentVisitor{},
		&AttributeVisitor{},
		&ReservedVisitor{},
	)
	return out
}

func (ov *OneOfVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "oneof ") && in.Token == OpenBrace
}

func (ov *OneOfVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	Log.Debug("Visiting OneOf")
	values := in.SplitSyntax()
	name := values[1]

	out := NewOneOf(Join(Period, namespace, name), name, in.Comment)

	var comment = Comment("")

	for scanner.Scan() {
		line := scanner.ReadLine()

		Log.Debugf("Current Line: `%s`\n", line)

		if strings.HasSuffix(line.Token, CloseBrace) {
			break
		}
		for _, visitor := range ov.Visitors {
			if visitor.CanVisit(line) {
				rt := visitor.Visit(
					scanner,
					line,
					Join(Period, namespace, out.Name))
				switch t := rt.(type) {
				case *Attribute:
					if t.IsValid() {
						t.Comment = comment.AddSpace().Append(t.Comment).TrimSpace()
						out.Attributes = append(out.Attributes, t)
						comment = comment.Clear()
					}
				case *Reserved:
					out.Reserved = append(out.Reserved, t)
				case Comment:
					comment = comment.Append(t).AddSpace()
				}
			}
		}
	}
	return out
}
