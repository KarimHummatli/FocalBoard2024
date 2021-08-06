// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
import React from 'react'
import {useIntl, IntlShape} from 'react-intl'

import {IContentBlockWithCords, ContentBlock as ContentBlockType} from '../../blocks/contentBlock'
import {Card} from '../../blocks/card'
import {createTextBlock} from '../../blocks/textBlock'
import mutator from '../../mutator'
import {useSortableWithGrip} from '../../hooks/sortable'

import ContentBlock from '../contentBlock'
import {MarkdownEditor} from '../markdownEditor'

export type Position = 'left' | 'right' | 'above' | 'below' | 'aboveRow' | 'belowRow'

type Props = {
    card: Card
    contents: Array<ContentBlockType|ContentBlockType[]>
    readonly: boolean
}

function addTextBlock(card: Card, intl: IntlShape, text: string): void {
    const block = createTextBlock()
    block.parentId = card.id
    block.rootId = card.rootId
    block.title = text

    const contentOrder = card.fields.contentOrder.slice()
    contentOrder.push(block.id)
    mutator.performAsUndoGroup(async () => {
        const description = intl.formatMessage({id: 'CardDetail.addCardText', defaultMessage: 'add card text'})
        await mutator.insertBlock(block, description)
        await mutator.changeCardContentOrder(card.id, card.fields.contentOrder, contentOrder, description)
    })
}

function moveBlock(card: Card, srcBlock: IContentBlockWithCords, dstBlock: IContentBlockWithCords, intl: IntlShape, moveTo: Position): void {
    const contentOrder: Array<string|string[]> = []
    if (card.fields.contentOrder) {
        for (const contentId of card.fields.contentOrder) {
            if (typeof contentId === 'string') {
                contentOrder.push(contentId)
            } else {
                contentOrder.push(contentId.slice())
            }
        }
    }

    const srcBlockId = srcBlock.block.id
    const dstBlockId = dstBlock.block.id

    const srcBlockX = srcBlock.cords.x
    let dstBlockX = dstBlock.cords.x

    const srcBlockY = (srcBlock.cords.y || srcBlock.cords.y === 0) && (srcBlock.cords.y > -1) ? srcBlock.cords.y : -1
    let dstBlockY = (dstBlock.cords.y || dstBlock.cords.y === 0) && (dstBlock.cords.y > -1) ? dstBlock.cords.y : -1

    if (srcBlockId === dstBlockId) {
        return
    }

    // Delete Src Block
    if (srcBlockY > -1) {
        (contentOrder[srcBlockX] as string[]).splice(srcBlockY, 1)

        if (contentOrder[srcBlockX].length === 1 && srcBlockX !== dstBlockX) {
            contentOrder.splice(srcBlockX, 1, contentOrder[srcBlockX][0])
        }
    } else {
        contentOrder.splice(srcBlockX, 1)

        if (dstBlockX > srcBlockX) {
            dstBlockX -= 1
        }
    }

    if (moveTo === 'right') {
        if (dstBlockY > -1) {
            if (dstBlockX === srcBlockX && dstBlockY > srcBlockY) {
                dstBlockY -= 1
            }

            (contentOrder[dstBlockX] as string[]).splice(dstBlockY + 1, 0, srcBlockId)
        } else {
            contentOrder.splice(dstBlockX, 1, [dstBlockId, srcBlockId])
        }
    } else if (moveTo === 'left') {
        if (dstBlockY > -1) {
            if (dstBlockX === srcBlockX && dstBlockY > srcBlockY) {
                dstBlockY -= 1
            }

            (contentOrder[dstBlockX] as string[]).splice(dstBlockY, 0, srcBlockId)
        } else {
            contentOrder.splice(dstBlockX, 1, [srcBlockId, dstBlockId])
        }
    } else if (moveTo === 'aboveRow') {
        contentOrder.splice(dstBlockX, 0, srcBlockId)
    } else if (moveTo === 'belowRow') {
        contentOrder.splice(dstBlockX + 1, 0, srcBlockId)
    }

    mutator.performAsUndoGroup(async () => {
        const description = intl.formatMessage({id: 'CardDetail.moveContent', defaultMessage: 'move card content'})
        await mutator.changeCardContentOrder(card.id, card.fields.contentOrder, contentOrder, description)
    })
}

type ContentBlockWithDragAndDropProps = {
    block: ContentBlockType | ContentBlockType[],
    x: number,
    card: Card,
    contents: Array<ContentBlockType|ContentBlockType[]>,
    intl: IntlShape,
    readonly: boolean,
}

const ContentBlockWithDragAndDrop = (props: ContentBlockWithDragAndDropProps) => {
    const [, isOver,, itemRef] = useSortableWithGrip('content', {block: props.block, cords: {x: props.x}}, true, (src, dst) => moveBlock(props.card, src, dst, props.intl, 'aboveRow'))
    const [, isOver2,, itemRef2] = useSortableWithGrip('content', {block: props.block, cords: {x: props.x}}, true, (src, dst) => moveBlock(props.card, src, dst, props.intl, 'belowRow'))

    if (Array.isArray(props.block)) {
        return (
            <div >
                <div
                    ref={itemRef}
                    className={`addToRow ${isOver ? 'dragover' : ''}`}
                    style={{width: '94%', height: '10px', marginLeft: '48px'}}
                />
                <div
                    style={{display: 'flex'}}
                >

                    {props.block.map((b, y) => (
                        <ContentBlock
                            key={b.id}
                            block={b}
                            card={props.card}
                            readonly={props.readonly}
                            width={(1 / (props.block as ContentBlockType[]).length) * 100}
                            onDrop={(src, dst, moveTo) => moveBlock(props.card, src, dst, props.intl, moveTo)}
                            cords={{x: props.x, y}}
                        />
                    ))}
                </div>
                {props.x === props.contents.length - 1 && (
                    <div
                        ref={itemRef2}
                        className={`addToRow ${isOver2 ? 'dragover' : ''}`}
                        style={{width: '94%', height: '10px', marginLeft: '48px'}}
                    />
                )}
            </div>

        )
    }

    return (
        <div>
            <div
                ref={itemRef}
                className={`addToRow ${isOver ? 'dragover' : ''}`}
                style={{width: '94%', height: '10px', marginLeft: '48px'}}
            />
            <ContentBlock
                key={props.block.id}
                block={props.block}
                card={props.card}
                readonly={props.readonly}
                onDrop={(src, dst, moveTo) => moveBlock(props.card, src, dst, props.intl, moveTo)}
                cords={{x: props.x}}
            />
            {props.x === props.contents.length - 1 && (
                <div
                    ref={itemRef2}
                    className={`addToRow ${isOver2 ? 'dragover' : ''}`}
                    style={{width: '94%', height: '10px', marginLeft: '48px'}}
                />
            )}
        </div>

    )
}

const CardDetailContents = React.memo((props: Props) => {
    const intl = useIntl()
    const {contents, card} = props
    if (!contents) {
        return null
    }
    if (contents.length > 0) {
        return (
            <div className='octo-content'>
                {contents.map((block, x) =>
                    (
                        <ContentBlockWithDragAndDrop
                            key={x}
                            block={block}
                            x={x}
                            card={card}
                            contents={contents}
                            intl={intl}
                            readonly={props.readonly}
                        />
                    ),
                )}
            </div>
        )
    }
    return (
        <div className='octo-content CardDetailContents'>
            <div className='octo-block'>
                <div className='octo-block-margin'/>
                {!props.readonly &&
                    <MarkdownEditor
                        text=''
                        placeholderText='Add a description...'
                        onBlur={(text) => {
                            if (text) {
                                addTextBlock(card, intl, text)
                            }
                        }}
                    />
                }
            </div>
        </div>
    )
})

export default CardDetailContents
