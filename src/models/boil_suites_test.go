// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Brands", testBrands)
	t.Run("BroadcastEntries", testBroadcastEntries)
	t.Run("BroadcastNotices", testBroadcastNotices)
	t.Run("CertificateSessions", testCertificateSessions)
	t.Run("Clients", testClients)
	t.Run("ClientAllowRules", testClientAllowRules)
	t.Run("ClientQuizzes", testClientQuizzes)
	t.Run("ClientRefreshes", testClientRefreshes)
	t.Run("ClientScopes", testClientScopes)
	t.Run("ClientSessions", testClientSessions)
	t.Run("EmailVerifySessions", testEmailVerifySessions)
	t.Run("LoginClients", testLoginClients)
	t.Run("LoginClientHistories", testLoginClientHistories)
	t.Run("LoginClientScopes", testLoginClientScopes)
	t.Run("LoginHistories", testLoginHistories)
	t.Run("LoginTryHistories", testLoginTryHistories)
	t.Run("OauthSessions", testOauthSessions)
	t.Run("Otps", testOtps)
	t.Run("OtpBackups", testOtpBackups)
	t.Run("OtpSessions", testOtpSessions)
	t.Run("Passkeys", testPasskeys)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevices)
	t.Run("Passwords", testPasswords)
	t.Run("Refreshes", testRefreshes)
	t.Run("RegisterOtpSessions", testRegisterOtpSessions)
	t.Run("RegisterSessions", testRegisterSessions)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessions)
	t.Run("Sessions", testSessions)
	t.Run("Settings", testSettings)
	t.Run("Staffs", testStaffs)
	t.Run("Users", testUsers)
	t.Run("WebauthnSessions", testWebauthnSessions)
}

func TestDelete(t *testing.T) {
	t.Run("Brands", testBrandsDelete)
	t.Run("BroadcastEntries", testBroadcastEntriesDelete)
	t.Run("BroadcastNotices", testBroadcastNoticesDelete)
	t.Run("CertificateSessions", testCertificateSessionsDelete)
	t.Run("Clients", testClientsDelete)
	t.Run("ClientAllowRules", testClientAllowRulesDelete)
	t.Run("ClientQuizzes", testClientQuizzesDelete)
	t.Run("ClientRefreshes", testClientRefreshesDelete)
	t.Run("ClientScopes", testClientScopesDelete)
	t.Run("ClientSessions", testClientSessionsDelete)
	t.Run("EmailVerifySessions", testEmailVerifySessionsDelete)
	t.Run("LoginClients", testLoginClientsDelete)
	t.Run("LoginClientHistories", testLoginClientHistoriesDelete)
	t.Run("LoginClientScopes", testLoginClientScopesDelete)
	t.Run("LoginHistories", testLoginHistoriesDelete)
	t.Run("LoginTryHistories", testLoginTryHistoriesDelete)
	t.Run("OauthSessions", testOauthSessionsDelete)
	t.Run("Otps", testOtpsDelete)
	t.Run("OtpBackups", testOtpBackupsDelete)
	t.Run("OtpSessions", testOtpSessionsDelete)
	t.Run("Passkeys", testPasskeysDelete)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesDelete)
	t.Run("Passwords", testPasswordsDelete)
	t.Run("Refreshes", testRefreshesDelete)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsDelete)
	t.Run("RegisterSessions", testRegisterSessionsDelete)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsDelete)
	t.Run("Sessions", testSessionsDelete)
	t.Run("Settings", testSettingsDelete)
	t.Run("Staffs", testStaffsDelete)
	t.Run("Users", testUsersDelete)
	t.Run("WebauthnSessions", testWebauthnSessionsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Brands", testBrandsQueryDeleteAll)
	t.Run("BroadcastEntries", testBroadcastEntriesQueryDeleteAll)
	t.Run("BroadcastNotices", testBroadcastNoticesQueryDeleteAll)
	t.Run("CertificateSessions", testCertificateSessionsQueryDeleteAll)
	t.Run("Clients", testClientsQueryDeleteAll)
	t.Run("ClientAllowRules", testClientAllowRulesQueryDeleteAll)
	t.Run("ClientQuizzes", testClientQuizzesQueryDeleteAll)
	t.Run("ClientRefreshes", testClientRefreshesQueryDeleteAll)
	t.Run("ClientScopes", testClientScopesQueryDeleteAll)
	t.Run("ClientSessions", testClientSessionsQueryDeleteAll)
	t.Run("EmailVerifySessions", testEmailVerifySessionsQueryDeleteAll)
	t.Run("LoginClients", testLoginClientsQueryDeleteAll)
	t.Run("LoginClientHistories", testLoginClientHistoriesQueryDeleteAll)
	t.Run("LoginClientScopes", testLoginClientScopesQueryDeleteAll)
	t.Run("LoginHistories", testLoginHistoriesQueryDeleteAll)
	t.Run("LoginTryHistories", testLoginTryHistoriesQueryDeleteAll)
	t.Run("OauthSessions", testOauthSessionsQueryDeleteAll)
	t.Run("Otps", testOtpsQueryDeleteAll)
	t.Run("OtpBackups", testOtpBackupsQueryDeleteAll)
	t.Run("OtpSessions", testOtpSessionsQueryDeleteAll)
	t.Run("Passkeys", testPasskeysQueryDeleteAll)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesQueryDeleteAll)
	t.Run("Passwords", testPasswordsQueryDeleteAll)
	t.Run("Refreshes", testRefreshesQueryDeleteAll)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsQueryDeleteAll)
	t.Run("RegisterSessions", testRegisterSessionsQueryDeleteAll)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsQueryDeleteAll)
	t.Run("Sessions", testSessionsQueryDeleteAll)
	t.Run("Settings", testSettingsQueryDeleteAll)
	t.Run("Staffs", testStaffsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("WebauthnSessions", testWebauthnSessionsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Brands", testBrandsSliceDeleteAll)
	t.Run("BroadcastEntries", testBroadcastEntriesSliceDeleteAll)
	t.Run("BroadcastNotices", testBroadcastNoticesSliceDeleteAll)
	t.Run("CertificateSessions", testCertificateSessionsSliceDeleteAll)
	t.Run("Clients", testClientsSliceDeleteAll)
	t.Run("ClientAllowRules", testClientAllowRulesSliceDeleteAll)
	t.Run("ClientQuizzes", testClientQuizzesSliceDeleteAll)
	t.Run("ClientRefreshes", testClientRefreshesSliceDeleteAll)
	t.Run("ClientScopes", testClientScopesSliceDeleteAll)
	t.Run("ClientSessions", testClientSessionsSliceDeleteAll)
	t.Run("EmailVerifySessions", testEmailVerifySessionsSliceDeleteAll)
	t.Run("LoginClients", testLoginClientsSliceDeleteAll)
	t.Run("LoginClientHistories", testLoginClientHistoriesSliceDeleteAll)
	t.Run("LoginClientScopes", testLoginClientScopesSliceDeleteAll)
	t.Run("LoginHistories", testLoginHistoriesSliceDeleteAll)
	t.Run("LoginTryHistories", testLoginTryHistoriesSliceDeleteAll)
	t.Run("OauthSessions", testOauthSessionsSliceDeleteAll)
	t.Run("Otps", testOtpsSliceDeleteAll)
	t.Run("OtpBackups", testOtpBackupsSliceDeleteAll)
	t.Run("OtpSessions", testOtpSessionsSliceDeleteAll)
	t.Run("Passkeys", testPasskeysSliceDeleteAll)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesSliceDeleteAll)
	t.Run("Passwords", testPasswordsSliceDeleteAll)
	t.Run("Refreshes", testRefreshesSliceDeleteAll)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsSliceDeleteAll)
	t.Run("RegisterSessions", testRegisterSessionsSliceDeleteAll)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsSliceDeleteAll)
	t.Run("Sessions", testSessionsSliceDeleteAll)
	t.Run("Settings", testSettingsSliceDeleteAll)
	t.Run("Staffs", testStaffsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("WebauthnSessions", testWebauthnSessionsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Brands", testBrandsExists)
	t.Run("BroadcastEntries", testBroadcastEntriesExists)
	t.Run("BroadcastNotices", testBroadcastNoticesExists)
	t.Run("CertificateSessions", testCertificateSessionsExists)
	t.Run("Clients", testClientsExists)
	t.Run("ClientAllowRules", testClientAllowRulesExists)
	t.Run("ClientQuizzes", testClientQuizzesExists)
	t.Run("ClientRefreshes", testClientRefreshesExists)
	t.Run("ClientScopes", testClientScopesExists)
	t.Run("ClientSessions", testClientSessionsExists)
	t.Run("EmailVerifySessions", testEmailVerifySessionsExists)
	t.Run("LoginClients", testLoginClientsExists)
	t.Run("LoginClientHistories", testLoginClientHistoriesExists)
	t.Run("LoginClientScopes", testLoginClientScopesExists)
	t.Run("LoginHistories", testLoginHistoriesExists)
	t.Run("LoginTryHistories", testLoginTryHistoriesExists)
	t.Run("OauthSessions", testOauthSessionsExists)
	t.Run("Otps", testOtpsExists)
	t.Run("OtpBackups", testOtpBackupsExists)
	t.Run("OtpSessions", testOtpSessionsExists)
	t.Run("Passkeys", testPasskeysExists)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesExists)
	t.Run("Passwords", testPasswordsExists)
	t.Run("Refreshes", testRefreshesExists)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsExists)
	t.Run("RegisterSessions", testRegisterSessionsExists)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsExists)
	t.Run("Sessions", testSessionsExists)
	t.Run("Settings", testSettingsExists)
	t.Run("Staffs", testStaffsExists)
	t.Run("Users", testUsersExists)
	t.Run("WebauthnSessions", testWebauthnSessionsExists)
}

func TestFind(t *testing.T) {
	t.Run("Brands", testBrandsFind)
	t.Run("BroadcastEntries", testBroadcastEntriesFind)
	t.Run("BroadcastNotices", testBroadcastNoticesFind)
	t.Run("CertificateSessions", testCertificateSessionsFind)
	t.Run("Clients", testClientsFind)
	t.Run("ClientAllowRules", testClientAllowRulesFind)
	t.Run("ClientQuizzes", testClientQuizzesFind)
	t.Run("ClientRefreshes", testClientRefreshesFind)
	t.Run("ClientScopes", testClientScopesFind)
	t.Run("ClientSessions", testClientSessionsFind)
	t.Run("EmailVerifySessions", testEmailVerifySessionsFind)
	t.Run("LoginClients", testLoginClientsFind)
	t.Run("LoginClientHistories", testLoginClientHistoriesFind)
	t.Run("LoginClientScopes", testLoginClientScopesFind)
	t.Run("LoginHistories", testLoginHistoriesFind)
	t.Run("LoginTryHistories", testLoginTryHistoriesFind)
	t.Run("OauthSessions", testOauthSessionsFind)
	t.Run("Otps", testOtpsFind)
	t.Run("OtpBackups", testOtpBackupsFind)
	t.Run("OtpSessions", testOtpSessionsFind)
	t.Run("Passkeys", testPasskeysFind)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesFind)
	t.Run("Passwords", testPasswordsFind)
	t.Run("Refreshes", testRefreshesFind)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsFind)
	t.Run("RegisterSessions", testRegisterSessionsFind)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsFind)
	t.Run("Sessions", testSessionsFind)
	t.Run("Settings", testSettingsFind)
	t.Run("Staffs", testStaffsFind)
	t.Run("Users", testUsersFind)
	t.Run("WebauthnSessions", testWebauthnSessionsFind)
}

func TestBind(t *testing.T) {
	t.Run("Brands", testBrandsBind)
	t.Run("BroadcastEntries", testBroadcastEntriesBind)
	t.Run("BroadcastNotices", testBroadcastNoticesBind)
	t.Run("CertificateSessions", testCertificateSessionsBind)
	t.Run("Clients", testClientsBind)
	t.Run("ClientAllowRules", testClientAllowRulesBind)
	t.Run("ClientQuizzes", testClientQuizzesBind)
	t.Run("ClientRefreshes", testClientRefreshesBind)
	t.Run("ClientScopes", testClientScopesBind)
	t.Run("ClientSessions", testClientSessionsBind)
	t.Run("EmailVerifySessions", testEmailVerifySessionsBind)
	t.Run("LoginClients", testLoginClientsBind)
	t.Run("LoginClientHistories", testLoginClientHistoriesBind)
	t.Run("LoginClientScopes", testLoginClientScopesBind)
	t.Run("LoginHistories", testLoginHistoriesBind)
	t.Run("LoginTryHistories", testLoginTryHistoriesBind)
	t.Run("OauthSessions", testOauthSessionsBind)
	t.Run("Otps", testOtpsBind)
	t.Run("OtpBackups", testOtpBackupsBind)
	t.Run("OtpSessions", testOtpSessionsBind)
	t.Run("Passkeys", testPasskeysBind)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesBind)
	t.Run("Passwords", testPasswordsBind)
	t.Run("Refreshes", testRefreshesBind)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsBind)
	t.Run("RegisterSessions", testRegisterSessionsBind)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsBind)
	t.Run("Sessions", testSessionsBind)
	t.Run("Settings", testSettingsBind)
	t.Run("Staffs", testStaffsBind)
	t.Run("Users", testUsersBind)
	t.Run("WebauthnSessions", testWebauthnSessionsBind)
}

func TestOne(t *testing.T) {
	t.Run("Brands", testBrandsOne)
	t.Run("BroadcastEntries", testBroadcastEntriesOne)
	t.Run("BroadcastNotices", testBroadcastNoticesOne)
	t.Run("CertificateSessions", testCertificateSessionsOne)
	t.Run("Clients", testClientsOne)
	t.Run("ClientAllowRules", testClientAllowRulesOne)
	t.Run("ClientQuizzes", testClientQuizzesOne)
	t.Run("ClientRefreshes", testClientRefreshesOne)
	t.Run("ClientScopes", testClientScopesOne)
	t.Run("ClientSessions", testClientSessionsOne)
	t.Run("EmailVerifySessions", testEmailVerifySessionsOne)
	t.Run("LoginClients", testLoginClientsOne)
	t.Run("LoginClientHistories", testLoginClientHistoriesOne)
	t.Run("LoginClientScopes", testLoginClientScopesOne)
	t.Run("LoginHistories", testLoginHistoriesOne)
	t.Run("LoginTryHistories", testLoginTryHistoriesOne)
	t.Run("OauthSessions", testOauthSessionsOne)
	t.Run("Otps", testOtpsOne)
	t.Run("OtpBackups", testOtpBackupsOne)
	t.Run("OtpSessions", testOtpSessionsOne)
	t.Run("Passkeys", testPasskeysOne)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesOne)
	t.Run("Passwords", testPasswordsOne)
	t.Run("Refreshes", testRefreshesOne)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsOne)
	t.Run("RegisterSessions", testRegisterSessionsOne)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsOne)
	t.Run("Sessions", testSessionsOne)
	t.Run("Settings", testSettingsOne)
	t.Run("Staffs", testStaffsOne)
	t.Run("Users", testUsersOne)
	t.Run("WebauthnSessions", testWebauthnSessionsOne)
}

func TestAll(t *testing.T) {
	t.Run("Brands", testBrandsAll)
	t.Run("BroadcastEntries", testBroadcastEntriesAll)
	t.Run("BroadcastNotices", testBroadcastNoticesAll)
	t.Run("CertificateSessions", testCertificateSessionsAll)
	t.Run("Clients", testClientsAll)
	t.Run("ClientAllowRules", testClientAllowRulesAll)
	t.Run("ClientQuizzes", testClientQuizzesAll)
	t.Run("ClientRefreshes", testClientRefreshesAll)
	t.Run("ClientScopes", testClientScopesAll)
	t.Run("ClientSessions", testClientSessionsAll)
	t.Run("EmailVerifySessions", testEmailVerifySessionsAll)
	t.Run("LoginClients", testLoginClientsAll)
	t.Run("LoginClientHistories", testLoginClientHistoriesAll)
	t.Run("LoginClientScopes", testLoginClientScopesAll)
	t.Run("LoginHistories", testLoginHistoriesAll)
	t.Run("LoginTryHistories", testLoginTryHistoriesAll)
	t.Run("OauthSessions", testOauthSessionsAll)
	t.Run("Otps", testOtpsAll)
	t.Run("OtpBackups", testOtpBackupsAll)
	t.Run("OtpSessions", testOtpSessionsAll)
	t.Run("Passkeys", testPasskeysAll)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesAll)
	t.Run("Passwords", testPasswordsAll)
	t.Run("Refreshes", testRefreshesAll)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsAll)
	t.Run("RegisterSessions", testRegisterSessionsAll)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsAll)
	t.Run("Sessions", testSessionsAll)
	t.Run("Settings", testSettingsAll)
	t.Run("Staffs", testStaffsAll)
	t.Run("Users", testUsersAll)
	t.Run("WebauthnSessions", testWebauthnSessionsAll)
}

func TestCount(t *testing.T) {
	t.Run("Brands", testBrandsCount)
	t.Run("BroadcastEntries", testBroadcastEntriesCount)
	t.Run("BroadcastNotices", testBroadcastNoticesCount)
	t.Run("CertificateSessions", testCertificateSessionsCount)
	t.Run("Clients", testClientsCount)
	t.Run("ClientAllowRules", testClientAllowRulesCount)
	t.Run("ClientQuizzes", testClientQuizzesCount)
	t.Run("ClientRefreshes", testClientRefreshesCount)
	t.Run("ClientScopes", testClientScopesCount)
	t.Run("ClientSessions", testClientSessionsCount)
	t.Run("EmailVerifySessions", testEmailVerifySessionsCount)
	t.Run("LoginClients", testLoginClientsCount)
	t.Run("LoginClientHistories", testLoginClientHistoriesCount)
	t.Run("LoginClientScopes", testLoginClientScopesCount)
	t.Run("LoginHistories", testLoginHistoriesCount)
	t.Run("LoginTryHistories", testLoginTryHistoriesCount)
	t.Run("OauthSessions", testOauthSessionsCount)
	t.Run("Otps", testOtpsCount)
	t.Run("OtpBackups", testOtpBackupsCount)
	t.Run("OtpSessions", testOtpSessionsCount)
	t.Run("Passkeys", testPasskeysCount)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesCount)
	t.Run("Passwords", testPasswordsCount)
	t.Run("Refreshes", testRefreshesCount)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsCount)
	t.Run("RegisterSessions", testRegisterSessionsCount)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsCount)
	t.Run("Sessions", testSessionsCount)
	t.Run("Settings", testSettingsCount)
	t.Run("Staffs", testStaffsCount)
	t.Run("Users", testUsersCount)
	t.Run("WebauthnSessions", testWebauthnSessionsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Brands", testBrandsHooks)
	t.Run("BroadcastEntries", testBroadcastEntriesHooks)
	t.Run("BroadcastNotices", testBroadcastNoticesHooks)
	t.Run("CertificateSessions", testCertificateSessionsHooks)
	t.Run("Clients", testClientsHooks)
	t.Run("ClientAllowRules", testClientAllowRulesHooks)
	t.Run("ClientQuizzes", testClientQuizzesHooks)
	t.Run("ClientRefreshes", testClientRefreshesHooks)
	t.Run("ClientScopes", testClientScopesHooks)
	t.Run("ClientSessions", testClientSessionsHooks)
	t.Run("EmailVerifySessions", testEmailVerifySessionsHooks)
	t.Run("LoginClients", testLoginClientsHooks)
	t.Run("LoginClientHistories", testLoginClientHistoriesHooks)
	t.Run("LoginClientScopes", testLoginClientScopesHooks)
	t.Run("LoginHistories", testLoginHistoriesHooks)
	t.Run("LoginTryHistories", testLoginTryHistoriesHooks)
	t.Run("OauthSessions", testOauthSessionsHooks)
	t.Run("Otps", testOtpsHooks)
	t.Run("OtpBackups", testOtpBackupsHooks)
	t.Run("OtpSessions", testOtpSessionsHooks)
	t.Run("Passkeys", testPasskeysHooks)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesHooks)
	t.Run("Passwords", testPasswordsHooks)
	t.Run("Refreshes", testRefreshesHooks)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsHooks)
	t.Run("RegisterSessions", testRegisterSessionsHooks)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsHooks)
	t.Run("Sessions", testSessionsHooks)
	t.Run("Settings", testSettingsHooks)
	t.Run("Staffs", testStaffsHooks)
	t.Run("Users", testUsersHooks)
	t.Run("WebauthnSessions", testWebauthnSessionsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Brands", testBrandsInsert)
	t.Run("Brands", testBrandsInsertWhitelist)
	t.Run("BroadcastEntries", testBroadcastEntriesInsert)
	t.Run("BroadcastEntries", testBroadcastEntriesInsertWhitelist)
	t.Run("BroadcastNotices", testBroadcastNoticesInsert)
	t.Run("BroadcastNotices", testBroadcastNoticesInsertWhitelist)
	t.Run("CertificateSessions", testCertificateSessionsInsert)
	t.Run("CertificateSessions", testCertificateSessionsInsertWhitelist)
	t.Run("Clients", testClientsInsert)
	t.Run("Clients", testClientsInsertWhitelist)
	t.Run("ClientAllowRules", testClientAllowRulesInsert)
	t.Run("ClientAllowRules", testClientAllowRulesInsertWhitelist)
	t.Run("ClientQuizzes", testClientQuizzesInsert)
	t.Run("ClientQuizzes", testClientQuizzesInsertWhitelist)
	t.Run("ClientRefreshes", testClientRefreshesInsert)
	t.Run("ClientRefreshes", testClientRefreshesInsertWhitelist)
	t.Run("ClientScopes", testClientScopesInsert)
	t.Run("ClientScopes", testClientScopesInsertWhitelist)
	t.Run("ClientSessions", testClientSessionsInsert)
	t.Run("ClientSessions", testClientSessionsInsertWhitelist)
	t.Run("EmailVerifySessions", testEmailVerifySessionsInsert)
	t.Run("EmailVerifySessions", testEmailVerifySessionsInsertWhitelist)
	t.Run("LoginClients", testLoginClientsInsert)
	t.Run("LoginClients", testLoginClientsInsertWhitelist)
	t.Run("LoginClientHistories", testLoginClientHistoriesInsert)
	t.Run("LoginClientHistories", testLoginClientHistoriesInsertWhitelist)
	t.Run("LoginClientScopes", testLoginClientScopesInsert)
	t.Run("LoginClientScopes", testLoginClientScopesInsertWhitelist)
	t.Run("LoginHistories", testLoginHistoriesInsert)
	t.Run("LoginHistories", testLoginHistoriesInsertWhitelist)
	t.Run("LoginTryHistories", testLoginTryHistoriesInsert)
	t.Run("LoginTryHistories", testLoginTryHistoriesInsertWhitelist)
	t.Run("OauthSessions", testOauthSessionsInsert)
	t.Run("OauthSessions", testOauthSessionsInsertWhitelist)
	t.Run("Otps", testOtpsInsert)
	t.Run("Otps", testOtpsInsertWhitelist)
	t.Run("OtpBackups", testOtpBackupsInsert)
	t.Run("OtpBackups", testOtpBackupsInsertWhitelist)
	t.Run("OtpSessions", testOtpSessionsInsert)
	t.Run("OtpSessions", testOtpSessionsInsertWhitelist)
	t.Run("Passkeys", testPasskeysInsert)
	t.Run("Passkeys", testPasskeysInsertWhitelist)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesInsert)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesInsertWhitelist)
	t.Run("Passwords", testPasswordsInsert)
	t.Run("Passwords", testPasswordsInsertWhitelist)
	t.Run("Refreshes", testRefreshesInsert)
	t.Run("Refreshes", testRefreshesInsertWhitelist)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsInsert)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsInsertWhitelist)
	t.Run("RegisterSessions", testRegisterSessionsInsert)
	t.Run("RegisterSessions", testRegisterSessionsInsertWhitelist)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsInsert)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsInsertWhitelist)
	t.Run("Sessions", testSessionsInsert)
	t.Run("Sessions", testSessionsInsertWhitelist)
	t.Run("Settings", testSettingsInsert)
	t.Run("Settings", testSettingsInsertWhitelist)
	t.Run("Staffs", testStaffsInsert)
	t.Run("Staffs", testStaffsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("WebauthnSessions", testWebauthnSessionsInsert)
	t.Run("WebauthnSessions", testWebauthnSessionsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Brands", testBrandsReload)
	t.Run("BroadcastEntries", testBroadcastEntriesReload)
	t.Run("BroadcastNotices", testBroadcastNoticesReload)
	t.Run("CertificateSessions", testCertificateSessionsReload)
	t.Run("Clients", testClientsReload)
	t.Run("ClientAllowRules", testClientAllowRulesReload)
	t.Run("ClientQuizzes", testClientQuizzesReload)
	t.Run("ClientRefreshes", testClientRefreshesReload)
	t.Run("ClientScopes", testClientScopesReload)
	t.Run("ClientSessions", testClientSessionsReload)
	t.Run("EmailVerifySessions", testEmailVerifySessionsReload)
	t.Run("LoginClients", testLoginClientsReload)
	t.Run("LoginClientHistories", testLoginClientHistoriesReload)
	t.Run("LoginClientScopes", testLoginClientScopesReload)
	t.Run("LoginHistories", testLoginHistoriesReload)
	t.Run("LoginTryHistories", testLoginTryHistoriesReload)
	t.Run("OauthSessions", testOauthSessionsReload)
	t.Run("Otps", testOtpsReload)
	t.Run("OtpBackups", testOtpBackupsReload)
	t.Run("OtpSessions", testOtpSessionsReload)
	t.Run("Passkeys", testPasskeysReload)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesReload)
	t.Run("Passwords", testPasswordsReload)
	t.Run("Refreshes", testRefreshesReload)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsReload)
	t.Run("RegisterSessions", testRegisterSessionsReload)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsReload)
	t.Run("Sessions", testSessionsReload)
	t.Run("Settings", testSettingsReload)
	t.Run("Staffs", testStaffsReload)
	t.Run("Users", testUsersReload)
	t.Run("WebauthnSessions", testWebauthnSessionsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Brands", testBrandsReloadAll)
	t.Run("BroadcastEntries", testBroadcastEntriesReloadAll)
	t.Run("BroadcastNotices", testBroadcastNoticesReloadAll)
	t.Run("CertificateSessions", testCertificateSessionsReloadAll)
	t.Run("Clients", testClientsReloadAll)
	t.Run("ClientAllowRules", testClientAllowRulesReloadAll)
	t.Run("ClientQuizzes", testClientQuizzesReloadAll)
	t.Run("ClientRefreshes", testClientRefreshesReloadAll)
	t.Run("ClientScopes", testClientScopesReloadAll)
	t.Run("ClientSessions", testClientSessionsReloadAll)
	t.Run("EmailVerifySessions", testEmailVerifySessionsReloadAll)
	t.Run("LoginClients", testLoginClientsReloadAll)
	t.Run("LoginClientHistories", testLoginClientHistoriesReloadAll)
	t.Run("LoginClientScopes", testLoginClientScopesReloadAll)
	t.Run("LoginHistories", testLoginHistoriesReloadAll)
	t.Run("LoginTryHistories", testLoginTryHistoriesReloadAll)
	t.Run("OauthSessions", testOauthSessionsReloadAll)
	t.Run("Otps", testOtpsReloadAll)
	t.Run("OtpBackups", testOtpBackupsReloadAll)
	t.Run("OtpSessions", testOtpSessionsReloadAll)
	t.Run("Passkeys", testPasskeysReloadAll)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesReloadAll)
	t.Run("Passwords", testPasswordsReloadAll)
	t.Run("Refreshes", testRefreshesReloadAll)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsReloadAll)
	t.Run("RegisterSessions", testRegisterSessionsReloadAll)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsReloadAll)
	t.Run("Sessions", testSessionsReloadAll)
	t.Run("Settings", testSettingsReloadAll)
	t.Run("Staffs", testStaffsReloadAll)
	t.Run("Users", testUsersReloadAll)
	t.Run("WebauthnSessions", testWebauthnSessionsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Brands", testBrandsSelect)
	t.Run("BroadcastEntries", testBroadcastEntriesSelect)
	t.Run("BroadcastNotices", testBroadcastNoticesSelect)
	t.Run("CertificateSessions", testCertificateSessionsSelect)
	t.Run("Clients", testClientsSelect)
	t.Run("ClientAllowRules", testClientAllowRulesSelect)
	t.Run("ClientQuizzes", testClientQuizzesSelect)
	t.Run("ClientRefreshes", testClientRefreshesSelect)
	t.Run("ClientScopes", testClientScopesSelect)
	t.Run("ClientSessions", testClientSessionsSelect)
	t.Run("EmailVerifySessions", testEmailVerifySessionsSelect)
	t.Run("LoginClients", testLoginClientsSelect)
	t.Run("LoginClientHistories", testLoginClientHistoriesSelect)
	t.Run("LoginClientScopes", testLoginClientScopesSelect)
	t.Run("LoginHistories", testLoginHistoriesSelect)
	t.Run("LoginTryHistories", testLoginTryHistoriesSelect)
	t.Run("OauthSessions", testOauthSessionsSelect)
	t.Run("Otps", testOtpsSelect)
	t.Run("OtpBackups", testOtpBackupsSelect)
	t.Run("OtpSessions", testOtpSessionsSelect)
	t.Run("Passkeys", testPasskeysSelect)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesSelect)
	t.Run("Passwords", testPasswordsSelect)
	t.Run("Refreshes", testRefreshesSelect)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsSelect)
	t.Run("RegisterSessions", testRegisterSessionsSelect)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsSelect)
	t.Run("Sessions", testSessionsSelect)
	t.Run("Settings", testSettingsSelect)
	t.Run("Staffs", testStaffsSelect)
	t.Run("Users", testUsersSelect)
	t.Run("WebauthnSessions", testWebauthnSessionsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Brands", testBrandsUpdate)
	t.Run("BroadcastEntries", testBroadcastEntriesUpdate)
	t.Run("BroadcastNotices", testBroadcastNoticesUpdate)
	t.Run("CertificateSessions", testCertificateSessionsUpdate)
	t.Run("Clients", testClientsUpdate)
	t.Run("ClientAllowRules", testClientAllowRulesUpdate)
	t.Run("ClientQuizzes", testClientQuizzesUpdate)
	t.Run("ClientRefreshes", testClientRefreshesUpdate)
	t.Run("ClientScopes", testClientScopesUpdate)
	t.Run("ClientSessions", testClientSessionsUpdate)
	t.Run("EmailVerifySessions", testEmailVerifySessionsUpdate)
	t.Run("LoginClients", testLoginClientsUpdate)
	t.Run("LoginClientHistories", testLoginClientHistoriesUpdate)
	t.Run("LoginClientScopes", testLoginClientScopesUpdate)
	t.Run("LoginHistories", testLoginHistoriesUpdate)
	t.Run("LoginTryHistories", testLoginTryHistoriesUpdate)
	t.Run("OauthSessions", testOauthSessionsUpdate)
	t.Run("Otps", testOtpsUpdate)
	t.Run("OtpBackups", testOtpBackupsUpdate)
	t.Run("OtpSessions", testOtpSessionsUpdate)
	t.Run("Passkeys", testPasskeysUpdate)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesUpdate)
	t.Run("Passwords", testPasswordsUpdate)
	t.Run("Refreshes", testRefreshesUpdate)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsUpdate)
	t.Run("RegisterSessions", testRegisterSessionsUpdate)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsUpdate)
	t.Run("Sessions", testSessionsUpdate)
	t.Run("Settings", testSettingsUpdate)
	t.Run("Staffs", testStaffsUpdate)
	t.Run("Users", testUsersUpdate)
	t.Run("WebauthnSessions", testWebauthnSessionsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Brands", testBrandsSliceUpdateAll)
	t.Run("BroadcastEntries", testBroadcastEntriesSliceUpdateAll)
	t.Run("BroadcastNotices", testBroadcastNoticesSliceUpdateAll)
	t.Run("CertificateSessions", testCertificateSessionsSliceUpdateAll)
	t.Run("Clients", testClientsSliceUpdateAll)
	t.Run("ClientAllowRules", testClientAllowRulesSliceUpdateAll)
	t.Run("ClientQuizzes", testClientQuizzesSliceUpdateAll)
	t.Run("ClientRefreshes", testClientRefreshesSliceUpdateAll)
	t.Run("ClientScopes", testClientScopesSliceUpdateAll)
	t.Run("ClientSessions", testClientSessionsSliceUpdateAll)
	t.Run("EmailVerifySessions", testEmailVerifySessionsSliceUpdateAll)
	t.Run("LoginClients", testLoginClientsSliceUpdateAll)
	t.Run("LoginClientHistories", testLoginClientHistoriesSliceUpdateAll)
	t.Run("LoginClientScopes", testLoginClientScopesSliceUpdateAll)
	t.Run("LoginHistories", testLoginHistoriesSliceUpdateAll)
	t.Run("LoginTryHistories", testLoginTryHistoriesSliceUpdateAll)
	t.Run("OauthSessions", testOauthSessionsSliceUpdateAll)
	t.Run("Otps", testOtpsSliceUpdateAll)
	t.Run("OtpBackups", testOtpBackupsSliceUpdateAll)
	t.Run("OtpSessions", testOtpSessionsSliceUpdateAll)
	t.Run("Passkeys", testPasskeysSliceUpdateAll)
	t.Run("PasskeyLoginDevices", testPasskeyLoginDevicesSliceUpdateAll)
	t.Run("Passwords", testPasswordsSliceUpdateAll)
	t.Run("Refreshes", testRefreshesSliceUpdateAll)
	t.Run("RegisterOtpSessions", testRegisterOtpSessionsSliceUpdateAll)
	t.Run("RegisterSessions", testRegisterSessionsSliceUpdateAll)
	t.Run("ReregistrationPasswordSessions", testReregistrationPasswordSessionsSliceUpdateAll)
	t.Run("Sessions", testSessionsSliceUpdateAll)
	t.Run("Settings", testSettingsSliceUpdateAll)
	t.Run("Staffs", testStaffsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("WebauthnSessions", testWebauthnSessionsSliceUpdateAll)
}
