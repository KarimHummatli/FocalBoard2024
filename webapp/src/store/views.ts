// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {createSlice, PayloadAction, createSelector} from '@reduxjs/toolkit'

import {BoardView, createBoardView} from '../blocks/boardView'
import {Utils} from '../utils'

import {initialReadOnlyLoad, loadBoardData} from './initialLoad'
import {getCurrentBoard} from './boards'

import {RootState} from './index'

type ViewsState = {
    current: string
    views: {[key: string]: BoardView}
}

const viewsSlice = createSlice({
    name: 'views',
    initialState: {views: {}, current: ''} as ViewsState,
    reducers: {
        setCurrent: (state, action: PayloadAction<string>) => {
            state.current = action.payload
        },
        updateViews: (state, action: PayloadAction<BoardView[]>) => {
            for (const view of action.payload) {
                if (view.deleteAt === 0) {
                    state.views[view.id] = view
                } else {
                    delete state.views[view.id]
                }
            }
        },
        updateView: (state, action: PayloadAction<BoardView>) => {
            state.views[action.payload.id] = action.payload
        },
    },
    extraReducers: (builder) => {
        builder.addCase(initialReadOnlyLoad.fulfilled, (state, action) => {
            state.views = {}
            for (const block of action.payload.blocks) {
                if (block.type === 'view') {
                    state.views[block.id] = block as BoardView
                }
            }
        })
        builder.addCase(loadBoardData.fulfilled, (state, action) => {
            state.views = {}
            for (const block of action.payload.blocks) {
                if (block.type === 'view') {
                    state.views[block.id] = block as BoardView
                }
            }
        })
    },
})

export const {updateViews, setCurrent, updateView} = viewsSlice.actions
export const {reducer} = viewsSlice

export const getViews = (state: RootState): {[key: string]: BoardView} => state.views.views
export const getSortedViews = createSelector(
    getViews,
    (views) => {
        return Object.values(views).sort((a, b) => a.title.localeCompare(b.title)).map((v) => createBoardView(v))
    },
)

export const getViewsByBoard = createSelector(
    getViews,
    (views) => {
        const result: {[key: string]: BoardView[]} = {}
        Object.values(views).forEach((view) => {
            if (result[view.parentId]) {
                result[view.parentId].push(view)
            } else {
                result[view.parentId] = [view]
            }
        })
        return result
    },
)

export function getView(viewId: string): (state: RootState) => BoardView|null {
    return (state: RootState): BoardView|null => {
        return state.views.views[viewId] || null
    }
}

export const getCurrentBoardViews = createSelector(
    (state) => state.boards.current,
    getViews,
    (boardId, views) => {
        Utils.log(`getCurrentBoardViews boardId: ${boardId} views: ${views.length}`)
        return Object.values(views).filter((v) => v.boardId === boardId).sort((a, b) => a.title.localeCompare(b.title)).map((v) => createBoardView(v))
    },
)

export const getCurrentView = createSelector(
    getViews,
    (state) => state.views.current,
    (views, viewId) => views[viewId],
)

export const getCurrentViewGroupBy = createSelector(
    getCurrentBoard,
    getCurrentView,
    (currentBoard, currentView) => {
        if (!currentBoard) {
            return undefined
        }
        if (!currentView) {
            return undefined
        }
        return currentBoard.cardProperties.find((o) => o.id === currentView.fields.groupById)
    },
)

export const getCurrentViewDisplayBy = createSelector(
    getCurrentBoard,
    getCurrentView,
    (currentBoard, currentView) => {
        if (!currentBoard) {
            return undefined
        }
        if (!currentView) {
            return undefined
        }
        return currentBoard.cardProperties.find((o) => o.id === currentView.fields.dateDisplayPropertyId)
    },
)
