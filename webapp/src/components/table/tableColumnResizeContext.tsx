// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
import React, {createContext, ReactElement, ReactNode, useCallback, useContext, useMemo} from 'react'
import {DragLayerMonitor, useDragLayer} from 'react-dnd'

import {Constants} from '../../constants'

export type ColumnResizeContextType = {
    updateRef: (cardId: string, columnId: string, element: HTMLDivElement | null) => void
    cellRef: (columnId: string) => HTMLDivElement | undefined
    width: (columnId: string) => number
}

const ColumnResizeContext = createContext<ColumnResizeContextType | null>(null)

export function useColumnResize(): ColumnResizeContextType {
    const context = useContext(ColumnResizeContext)
    if (!context) {
        throw new Error('ColumnResizeContext is not available!')
    }
    return context
}

export type ColumnResizeProviderProps = {
    children: ReactNode
    columnWidths: Record<string, number>
}

const columnWidth = (columnId: string, columnWidths: Record<string, number>, offset: number): string => {
    return `${Math.max(Constants.minColumnWidth, (columnWidths[columnId] || 0) + offset)}px`
}

export const ColumnResizeProvider = (props: ColumnResizeProviderProps): ReactElement => {
    const {children, columnWidths} = props

    type ElementsMap = Map<string, HTMLDivElement>
    const columns = useMemo(() => new Map<string, ElementsMap>(), [])

    const updateWidth = useCallback((columnId: string, elements: ElementsMap, offset: number) => {
        const width = columnWidth(columnId, columnWidths, offset)
        for (const element of elements.values()) {
            element.style.width = width
        }
    }, [columnWidths])

    const collect = useCallback((monitor: DragLayerMonitor<{id: string, width: number}>) => {
        if (monitor.getItemType() === 'horizontalGrip') {
            const difference = monitor.getDifferenceFromInitialOffset()
            if (difference) {
                const offset = difference.x
                const {id: columnId, width: itemWidth} = monitor.getItem()
                const width = columnWidths[columnId]
                const elements = columns.get(columnId)
                if (elements && itemWidth === width) {
                    updateWidth(columnId, elements, offset)
                }
                return offset
            }
        }
        return 0
    }, [updateWidth, columnWidths])

    useDragLayer(collect)

    const contextValue = useMemo((): ColumnResizeContextType => ({
        updateRef: (cardId, columnId, element) => {
            let elements = columns.get(columnId)
            if (element) {
                if (!elements) {
                    elements = new Map()
                    columns.set(columnId, elements)
                }
                elements.set(cardId, element)
            } else if (elements) {
                elements.delete(cardId)
            }
        },
        cellRef: (columnId): HTMLDivElement | undefined => {
            const iter = columns.get(columnId)?.values()
            if (iter) {
                const {value, done} = iter.next()
                return done ? value : iter.next().value
            }
            return undefined
        },
        width: (columnId) => {
            return Math.max(Constants.minColumnWidth, (columnWidths[columnId] || 0))
        },
    }), [columnWidths])

    return (
        <ColumnResizeContext.Provider value={contextValue}>
            {children}
        </ColumnResizeContext.Provider>
    )
}
