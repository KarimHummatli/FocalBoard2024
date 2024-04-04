// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
import React from 'react'
// import {useDrop} from 'react-dnd'
import { Droppable } from 'react-beautiful-dnd';

import {Card} from '../../blocks/card'
import './kanbanColumn.scss'

type Props = {
    // onDrop: (card: Card) => void
    children: React.ReactNode
    columnId: string; 
}
const KanbanColumn: React.FC<Props> = ({ children, columnId }) => {
    return (
        <Droppable droppableId={columnId}>
            {(provided) => (
                <div ref={provided.innerRef} 
                    {...provided.droppableProps} 
                    className="octo-board-column"
                >
                    <div 
                        className='octo-board-column-single' >
                            {children}
                            {provided.placeholder}
                    </div>
                </div>
            )}
        </Droppable>
    );
};

// const KanbanColumn = (props: Props) => {
//     const [{isOver}, drop] = useDrop(() => ({
//         accept: 'card',
//         collect: (monitor) => ({
//             isOver: monitor.isOver(),
//         }),
//         drop: (item: Card, monitor) => {
//             if (monitor.isOver({shallow: true})) {
//                 props.onDrop(item)
//             }
//         },
//     }), [props.onDrop])

//     let className = 'octo-board-column'
//     if (isOver) {
//         className += ' dragover'
//     }
//     return (
//         <div
//             ref={drop}
//             className={className}
//         >
//             <div className='octo-board-column-single'>
//                 {props.children}
//             </div>
//         </div>
//     )
// }

export default React.memo(KanbanColumn)
