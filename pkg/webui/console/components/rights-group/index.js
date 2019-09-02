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
import bind from 'autobind-decorator'
import { defineMessages, injectIntl } from 'react-intl'
import classnames from 'classnames'

import PropTypes from '../../../lib/prop-types'
import { RIGHT_ALL } from '../../lib/rights'

import Checkbox from '../../../components/checkbox'
import Notification from '../../../components/notification'
import Message from '../../../lib/components/message'

import style from './rights-group.styl'

const m = defineMessages({
  selectAll: 'Select All',
  outOfOwnScopeRights:
    'This {entityType} has more rights than you have. These rights can only be revoked.',
  outOfOwnScopeRightsStrict:
    'This {entityType} has more rights than you have. Modifying is hence prohibited.',
  outOfOwnScopeUniversalRight:
    "This {entityType} has a universal right that you don't have. Modifying the rights will revoke this right.",
  revocableOnly: 'revocable only',
})

const computeState = function(values, rights, universalRight) {
  const selectedCheckboxesCount = values.length
  const totalCheckboxesCount = rights.length

  const universalRightIsChecked = values.includes(RIGHT_ALL) || values.includes(universalRight)

  return {
    allSelected: universalRightIsChecked || selectedCheckboxesCount === totalCheckboxesCount,
    indeterminate:
      !universalRightIsChecked &&
      selectedCheckboxesCount !== 0 &&
      selectedCheckboxesCount !== totalCheckboxesCount,
  }
}

@injectIntl
@bind
class RightsGroup extends React.Component {
  static propTypes = {
    /** The class to be added to the container */
    className: PropTypes.string,
    /** A flag indicating whether the whole component should be disabled **/
    disabled: PropTypes.bool,
    /** The message depicting the type of entity this component is setting the
     * rights for.
     */
    entityTypeMessage: PropTypes.message.isRequired,
    /** The intl object provided by injectIntl of react-intl, used to translate
     * messages
     */
    intl: PropTypes.shape({
      formatMessage: PropTypes.func.isRequired,
    }).isRequired,
    /** The name prop, used to connect to formik */
    name: PropTypes.string.isRequired,
    /** Blur event hook */
    onBlur: PropTypes.func,
    /** Change event hook */
    onChange: PropTypes.func,
    /** The list of rights options */
    rights: PropTypes.rights,
    /** A flag identifying whether modifying rights is allowed when out of scope
     * rights are present. Can be used to prevent user error.
     */
    strict: PropTypes.bool,
    /** The universal right literal comprising all other rights */
    universalRight: PropTypes.string,
    /** The rights value */
    value: PropTypes.rights,
  }

  static defaultProps = {
    className: '',
    disabled: false,
    onBlur: () => null,
    onChange: () => null,
    rights: [],
    strict: false,
    universalRight: '',
    value: [],
  }

  state = {
    hasOutOfOwnScopeRights: undefined,
  }

  static getDerivedStateFromProps(props, state) {
    const { value, rights: grantableRights, universalRight: grantableUniversalRight } = props
    const universalRight =
      grantableUniversalRight || value.find(right => right !== RIGHT_ALL && right.endsWith('_ALL'))
    const allGrantableRights = [...grantableRights, ...grantableUniversalRight]
    let { outOfOwnScopeRights, hasOutOfOwnScopeRights, hasOutOfOwnScopeUniversalRight } = state

    // If the entity has more rights than user can set himself, these rights
    // will be present in the value prop, but not in the rights prop. Hence,
    // we need to extract these additional rights initially to be able to
    // display them persistently, even when they get unchecked
    if (hasOutOfOwnScopeRights === undefined) {
      // Filter out rights that the entity has but may not be granted by the
      // user
      outOfOwnScopeRights = value.filter(right => !allGrantableRights.includes(right))

      // Examine the result and also store whether universal rights are present
      hasOutOfOwnScopeRights = outOfOwnScopeRights.length !== 0
      // Store whether out of scope universal rights are present
      hasOutOfOwnScopeUniversalRight =
        outOfOwnScopeRights.filter(right => right.endsWith('_ALL')).length !== 0
      // Remove universal rights from the list
      outOfOwnScopeRights = outOfOwnScopeRights.filter(right => !right.endsWith('_ALL'))
    }

    // Compose rights list
    const rights = [...outOfOwnScopeRights, ...grantableRights]
    const { allSelected, indeterminate } = computeState(value, rights, universalRight)

    return {
      allSelected,
      indeterminate,
      outOfOwnScopeRights,
      hasOutOfOwnScopeRights,
      hasOutOfOwnScopeUniversalRight,
      rights,
      universalRight,
      value,
    }
  }

  async handleChangeAll(event) {
    const { onChange } = this.props
    const { rights, outOfOwnScopeRights, universalRight } = this.state
    const { checked } = event.target

    let value

    // Determine new value based on universal rights and out of scope rights
    if (checked) {
      if (universalRight) {
        // Prefer universal right value, if present
        value = [universalRight]
      } else {
        // Else add rights individually
        value = [...rights]
      }
    } else {
      // On uncheck, leave out of scope rights checked, if present
      value = [...outOfOwnScopeRights]
    }

    onChange(value)
  }

  async handleChange(val) {
    const value = Object.keys(val).filter(right => val[right])
    const { onChange, rights } = this.props
    const { universalRight } = this.state
    const { allSelected } = computeState(value, rights, universalRight)

    // Set new right value and prefer universal right if applicable
    const result = universalRight && allSelected ? [universalRight] : [...value]

    onChange(result)
  }

  render() {
    const { className, name, onBlur, disabled, strict, intl, entityTypeMessage } = this.props

    const {
      indeterminate,
      outOfOwnScopeRights,
      hasOutOfOwnScopeRights,
      hasOutOfOwnScopeUniversalRight,
      value,
      allSelected,
      rights,
      universalRight,
    } = this.state

    const hasOutOfOwnScopeNonUniversalRights =
      hasOutOfOwnScopeRights && !hasOutOfOwnScopeUniversalRight

    const cbs = rights.map(right => (
      <Checkbox
        className={style.rightLabel}
        key={right}
        name={right}
        disabled={outOfOwnScopeRights.includes(right)}
        label={{ id: `enum:${right}` }}
      >
        {outOfOwnScopeRights.includes(right) && (
          <Message className={style.revoke} content={m.revocableOnly} />
        )}
      </Checkbox>
    ))

    const allDisabled = disabled || (strict && hasOutOfOwnScopeNonUniversalRights)

    // Marshal rights to key/value for checkbox group
    const rightsValues = rights.reduce(function(acc, right) {
      acc[right] = allSelected || value.includes(right)

      return acc
    }, {})

    const collaboratorMessage = intl.formatMessage(entityTypeMessage).toLowerCase()

    return (
      <div className={className}>
        {hasOutOfOwnScopeUniversalRight && (
          <Notification
            small
            warning={m.outOfOwnScopeUniversalRight}
            messageValues={{ entityType: collaboratorMessage }}
          />
        )}
        {hasOutOfOwnScopeNonUniversalRights && (
          <Notification
            small
            warning={strict ? m.outOfOwnScopeRightsStrict : m.outOfOwnScopeRights}
            messageValues={{ entityType: collaboratorMessage }}
          />
        )}
        <Checkbox
          className={classnames(style.selectAll, style.rightLabel)}
          name={universalRight || 'select-all'}
          label={universalRight ? { id: `enum:${universalRight}` } : m.selectAll}
          onChange={this.handleChangeAll}
          indeterminate={indeterminate}
          value={allSelected}
          disabled={allDisabled}
        >
          {hasOutOfOwnScopeUniversalRight && universalRight && (
            <Message className={style.revoke} content={m.revocableOnly} />
          )}
        </Checkbox>
        <Checkbox.Group
          className={style.group}
          horizontal
          name={name}
          value={rightsValues}
          onChange={this.handleChange}
          onBlur={onBlur}
          disabled={allDisabled}
        >
          {cbs}
        </Checkbox.Group>
      </div>
    )
  }
}

export default RightsGroup
