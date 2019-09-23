// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import { defineMessages } from 'react-intl'
import bind from 'autobind-decorator'

import FetchTable from '../fetch-table'
import Message from '../../../lib/components/message'

import sharedMessages from '../../../lib/shared-messages'

import { getPubsubsList } from '../../../console/store/actions/pubsubs'

const m = defineMessages({
  format: 'Format',
  baseUrl: 'Base URL',
})

const headers = [
  {
    name: 'ids.pub_sub_id',
    displayName: sharedMessages.id,
    width: 35,
  },
  {
    name: 'format',
    displayName: m.format,
    width: 20,
  },
  {
    name: 'base_topic',
    displayName: m.baseUrl,
    width: 25,
  },
]

const getItemPathPrefix = item => `/${item.ids.pub_sub_id}`

@bind
export default class PubsubsTable extends React.Component {
  constructor(props) {
    super(props)

    const { appId } = props
    this.getPubsubsList = () => getPubsubsList(appId)
  }

  baseDataSelector({ pubsubs }) {
    return pubsubs
  }

  render() {
    return (
      <FetchTable
        entity="pubsubs"
        addMessage={sharedMessages.addPubsub}
        headers={headers}
        getItemsAction={this.getPubsubsList}
        baseDataSelector={this.baseDataSelector}
        tableTitle={<Message content={sharedMessages.pubsubs} />}
        getItemPathPrefix={getItemPathPrefix}
        {...this.props}
      />
    )
  }
}
