// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (c *Category) Todos(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *TodoOrder, where *TodoWhereInput,
) (*TodoConnection, error) {
	opts := []TodoPaginateOption{
		WithTodoOrder(orderBy),
		WithTodoFilter(where.Filter),
	}
	totalCount := c.Edges.totalCount[0]
	if nodes, err := c.Edges.TodosOrErr(); err == nil || totalCount != nil {
		pager, err := newTodoPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &TodoConnection{Edges: []*TodoEdge{}}
		if totalCount != nil {
			conn.TotalCount = *totalCount
		}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryTodos().Paginate(ctx, after, first, before, last, opts...)
}

func (f *Friendship) User(ctx context.Context) (*User, error) {
	result, err := f.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = f.QueryUser().Only(ctx)
	}
	return result, err
}

func (f *Friendship) Friend(ctx context.Context) (*User, error) {
	result, err := f.Edges.FriendOrErr()
	if IsNotLoaded(err) {
		result, err = f.QueryFriend().Only(ctx)
	}
	return result, err
}

func (gr *Group) Users(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, where *UserWhereInput,
) (*UserConnection, error) {
	opts := []UserPaginateOption{
		WithUserFilter(where.Filter),
	}
	totalCount := gr.Edges.totalCount[0]
	if nodes, err := gr.Edges.UsersOrErr(); err == nil || totalCount != nil {
		pager, err := newUserPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &UserConnection{Edges: []*UserEdge{}}
		if totalCount != nil {
			conn.TotalCount = *totalCount
		}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return gr.QueryUsers().Paginate(ctx, after, first, before, last, opts...)
}

func (t *Todo) Parent(ctx context.Context) (*Todo, error) {
	result, err := t.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Todo) Children(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *TodoOrder, where *TodoWhereInput,
) (*TodoConnection, error) {
	opts := []TodoPaginateOption{
		WithTodoOrder(orderBy),
		WithTodoFilter(where.Filter),
	}
	totalCount := t.Edges.totalCount[1]
	if nodes, err := t.Edges.ChildrenOrErr(); err == nil || totalCount != nil {
		pager, err := newTodoPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &TodoConnection{Edges: []*TodoEdge{}}
		if totalCount != nil {
			conn.TotalCount = *totalCount
		}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return t.QueryChildren().Paginate(ctx, after, first, before, last, opts...)
}

func (t *Todo) Category(ctx context.Context) (*Category, error) {
	result, err := t.Edges.CategoryOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryCategory().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Groups(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, where *GroupWhereInput,
) (*GroupConnection, error) {
	opts := []GroupPaginateOption{
		WithGroupFilter(where.Filter),
	}
	totalCount := u.Edges.totalCount[0]
	if nodes, err := u.Edges.GroupsOrErr(); err == nil || totalCount != nil {
		pager, err := newGroupPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &GroupConnection{Edges: []*GroupEdge{}}
		if totalCount != nil {
			conn.TotalCount = *totalCount
		}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryGroups().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) Friends(ctx context.Context) ([]*User, error) {
	result, err := u.Edges.FriendsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryFriends().All(ctx)
	}
	return result, err
}

func (u *User) Friendships(ctx context.Context) ([]*Friendship, error) {
	result, err := u.Edges.FriendshipsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryFriendships().All(ctx)
	}
	return result, err
}
