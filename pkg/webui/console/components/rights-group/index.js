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
import { defineMessages } from 'react-intl'
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
    'This entity possesses rights that are out of your scope of granted rights. These rights can only be revoked.',
  outOfOwnScopeRightsStrict:
    'This entity possesses rights that are out of your scope of granted rights. Modifying is hence prohibited.',
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

@bind
class RightsGroup extends React.Component {
  state = {
    hasOutOfOwnScopeRights: undefined,
  }

  static getDerivedStateFromProps(props, state) {
    const { value, rights: grantableRights, universalRight: grantableUniversalRight } = props
    let { outOfOwnScopeRights, hasOutOfOwnScopeRights, hasOutOfOwnScopeUniversalRight } = state
    const universalRight = grantableUniversalRight || value.find(right => right.endsWith('_ALL'))
    const allGrantableRights = [...grantableRights, ...grantableUniversalRight]

    // If the entity has more rights than user can set himself, these rights
    // will be present in the value prop, but not in the rights prop. Hence,
    // we need to extract these additional rights initially to be able to
    // display them persistently, even when they get unchecked
    if (hasOutOfOwnScopeRights === undefined) {
      // Filter out rights that the entity has but may not be granted by the
      // user
      outOfOwnScopeRights = value.filter(
        right => right !== RIGHT_ALL && !allGrantableRights.includes(right),
      )

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
    const { className, name, onBlur, disabled, strict } = this.props

    const {
      indeterminate,
      outOfOwnScopeRights,
      hasOutOfOwnScopeRights,
      hasOutOfOwnScopeUniversalRights,
      value,
      allSelected,
      rights,
      universalRight,
    } = this.state

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

    const allDisabled =
      disabled || (strict && hasOutOfOwnScopeRights) || hasOutOfOwnScopeUniversalRights

    // Marshal rights to key/value for checkbox group
    const rightsValues = rights.reduce(function(acc, right) {
      acc[right] = allSelected || value.includes(right)

      return acc
    }, {})

    return (
      <div className={className}>
        {hasOutOfOwnScopeRights && (
          <Notification small info={strict ? m.outOfOwnScopeRightsStrict : m.outOfOwnScopeRights} />
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
          {hasOutOfOwnScopeUniversalRights && (
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

RightsGroup.propTypes = {
  /** The class to be added to the container */
  className: PropTypes.string,
  /** The name prop, used to connect to formik */
  name: PropTypes.string.isRequired,
  /** The rights value */
  value: PropTypes.array,
  /** Change event hook */
  onChange: PropTypes.func,
  /** Blur event hook */
  onBlur: PropTypes.func,
  /** The universal right literal comprising all other rights */
  universalRight: PropTypes.string,
  /** The list of rights options */
  rights: PropTypes.arrayOf(PropTypes.string),
  /** A flag identifying whether modifying rights is allowed when out of scope
   * rights are present. Can be used to prevent user error.
   */
  strict: PropTypes.bool,
}

RightsGroup.defaultProps = {
  onChange: () => null,
  onBlur: () => null,
  rights: [],
  value: [],
  strict: false,
}

export default RightsGroup
