import React from 'react'

import { Unary } from '@berty/components/account'
import { useCreateNewAccount } from '@berty/hooks'
import { ScreenFC } from '@berty/navigation'
import { accountClient } from '@berty/utils/accounts/accountClient'

export const CreatingAccount: ScreenFC<'Account.Creating'> = () => {
	const createNewAccount = useCreateNewAccount()

	React.useEffect(() => {
		const f = async () => {
			// with an empty accountId the function returns default config
			const defaultConfig = await accountClient.networkConfigGet({ accountId: '' })
			if (defaultConfig.currentConfig) {
				await createNewAccount(defaultConfig.currentConfig)
			}
		}

		f()
	}, [createNewAccount])

	return <Unary>Creating Account...</Unary>
}
