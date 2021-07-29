// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
import {Utils} from '../utils'

import {Constants} from '../constants'

import {IBlock, MutableBlock} from './block'
import {IPropertyTemplate} from './board'

interface Card extends IBlock {
    readonly icon: string
    readonly isTemplate: boolean
    readonly properties: Readonly<Record<string, string | string[]>>
    readonly contentOrder: Readonly<Array<string | string[]>>

    duplicate(): MutableCard
    getProperty(property: IPropertyTemplate): string | string[] | number
}

class MutableCard extends MutableBlock implements Card {
    get icon(): string {
        return this.fields.icon as string
    }
    set icon(value: string) {
        this.fields.icon = value
    }

    get isTemplate(): boolean {
        return Boolean(this.fields.isTemplate)
    }
    set isTemplate(value: boolean) {
        this.fields.isTemplate = value
    }

    get properties(): Record<string, string | string[]> {
        return this.fields.properties as Record<string, string | string[]>
    }
    set properties(value: Record<string, string | string[]>) {
        this.fields.properties = value
    }

    get contentOrder(): Array<string | string[]> {
        return this.fields.contentOrder
    }
    set contentOrder(value: Array<string | string[]>) {
        this.fields.contentOrder = value
    }

    constructor(block: any = {}) {
        super(block)
        this.type = 'card'

        this.icon = block.fields?.icon || ''
        this.properties = {...(block.fields?.properties || {})}
        this.contentOrder = block.fields?.contentOrder?.slice() || []
    }

    duplicate(): MutableCard {
        const card = new MutableCard(this)
        card.id = Utils.createGuid()
        return card
    }

    getProperty(property: IPropertyTemplate): string | string[] | number {
        if (property.id === Constants.titleColumnId) {
            return this.title
        }

        switch (property.type) {
        case ('createdBy'): {
            return this.createdBy
        }
        case ('createdTime'): {
            return this.createAt
        }
        case ('updatedBy'): {
            return this.modifiedBy
        }
        case ('updatedTime'): {
            return this.updateAt
        }
        default: {
            return this.properties[property.id]
        }
        }
    }
}

export {MutableCard, Card}
