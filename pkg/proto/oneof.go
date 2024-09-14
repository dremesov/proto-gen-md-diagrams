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

// OneOf represents a Proto OneOf type.
type OneOf struct {
	*Qualified
	Attributes []*Attribute
	Reserved   []*Reserved
}

// NewOneOf is the OneOf Constructor
func NewOneOf(q string, name string, comment Comment) *OneOf {
	return &OneOf{
		Qualified: &Qualified{
			Qualifier: q,
			Name:      name,
			Comment:   comment,
		},
		Attributes: make([]*Attribute, 0),
		Reserved:   make([]*Reserved, 0),
	}
}

func (o *OneOf) ToMessage() *Message {
	return &Message{
		Qualified:  &Qualified{Qualifier: o.Qualifier, Name: o.Name, Comment: o.Comment},
		Attributes: o.Attributes,
		Messages:   []*Message{},
		OneOfs:     []*OneOf{},
		Enums:      []*Enum{},
		Reserved:   o.Reserved,
	}
}

// ToMermaid implements a Mermaid Syntax per Attribute
func (o *OneOf) ToMermaid() string {
	return Join(Space, "+", "OneOf", o.Name)
}
