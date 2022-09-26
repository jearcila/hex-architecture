package status

import (
	transactions_constants "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/constants"
	transactions_factory "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/factory"
)

var (
	Approved    = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_APPROVED}
	Contingency = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_CONTINGENCY}
	Scheduled   = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_SCHEDULED}

	Rejected                      = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_REJECTED}
	RejectedInvalidCardNumber     = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_INVALID_CARD_NUMBER}
	RejectedInvalidExpirationDate = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_INVALID_CARD_EXPIRATION_DATE}
	RejectedInvalidSecurityCode   = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_INVALID_CARD_SECURITY_CODE}
	RejectedCallForAuthorization  = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_CALL_FOR_AUTHORIZATION}
	RejectedInsufficientFunds     = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_INSUFFICIENT_LIMIT}
	RejectedDuplicatedPayment     = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_PAYMENT_DUPLICATED}
	RejectedBlackListed           = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_BLACKLISTED}
	RejectedFormError             = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_REJECTED_FORM_ERROR}
	RejectedMaxAttempts           = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_MAXIMUM_ATTEMPTS_REACHED}
	AmountRateLimitExceeded       = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_REJECTED_LIMIT_EXCEEDED}
	DisabledAccount               = transactions_factory.OperationStatus{Status: transactions_constants.STATUS_DISABLED_ACCOUNT}
)

func FindStatusByResponseCode(codeMap map[Code]transactions_factory.OperationStatus, responseCode string) transactions_factory.OperationStatus {
	statusInfo, ok := codeMap[responseCode]
	if !ok {
		statusInfo = Contingency
	}
	return statusInfo
}

var StatusByProviderCapture = map[Code]transactions_factory.OperationStatus{
	Code00:   Approved,
	CodeE001: Rejected,
	CodeE002: Contingency,
	CodeE004: Rejected,
	CodeR13:  Rejected,
}

var StatusByProviderRefund = map[Code]transactions_factory.OperationStatus{
	Code00:   Approved,
	Code01:   RejectedCallForAuthorization,
	Code02:   RejectedCallForAuthorization,
	Code03:   Rejected,
	Code04:   RejectedBlackListed,
	Code05:   Rejected,
	Code06:   RejectedFormError,
	Code07:   RejectedBlackListed,
	Code08:   Approved,
	Code10:   Approved,
	Code12:   Rejected,
	Code13:   RejectedFormError,
	Code14:   RejectedInvalidCardNumber,
	Code15:   Rejected,
	Code19:   Rejected,
	Code21:   Rejected,
	Code25:   Rejected,
	Code28:   Rejected,
	Code30:   Rejected,
	Code39:   RejectedFormError,
	Code41:   RejectedBlackListed,
	Code43:   RejectedBlackListed,
	Code46:   Rejected,
	Code51:   RejectedInsufficientFunds,
	Code52:   RejectedFormError,
	Code53:   RejectedFormError,
	Code54:   RejectedInvalidExpirationDate,
	Code55:   RejectedInvalidSecurityCode,
	Code57:   RejectedCallForAuthorization,
	Code58:   RejectedCallForAuthorization,
	Code59:   RejectedBlackListed,
	Code61:   RejectedInsufficientFunds,
	Code62:   RejectedCallForAuthorization,
	Code63:   RejectedCallForAuthorization,
	Code64:   RejectedBlackListed,
	Code65:   Rejected,
	Code6P:   RejectedFormError,
	Code70:   RejectedCallForAuthorization,
	Code71:   Rejected,
	Code74:   RejectedFormError,
	Code75:   Rejected,
	Code76:   Rejected,
	Code77:   Rejected,
	Code78:   DisabledAccount,
	Code79:   RejectedCallForAuthorization,
	Code81:   Rejected,
	Code82:   RejectedCallForAuthorization,
	Code83:   RejectedCallForAuthorization,
	Code84:   Rejected,
	Code85:   Approved,
	Code86:   RejectedInvalidSecurityCode,
	Code87:   Approved,
	Code88:   RejectedCallForAuthorization,
	Code89:   Rejected,
	Code91:   Rejected,
	Code92:   Rejected,
	Code93:   Rejected,
	Code94:   RejectedDuplicatedPayment,
	Code96:   Rejected,
	Code1A:   Rejected,
	CodeN0:   Scheduled,
	CodeN3:   Rejected,
	CodeN4:   Rejected,
	CodeN7:   RejectedInvalidSecurityCode,
	CodeN8:   Rejected,
	CodeQ1:   Rejected,
	CodeR0:   Rejected,
	CodeR1:   Rejected,
	CodeR2:   Rejected,
	CodeR3:   Rejected,
	CodeZ1:   Rejected,
	CodeZ3:   Rejected,
	CodeE001: Rejected,
	CodeE002: Scheduled,
	CodeE003: RejectedInvalidCardNumber,
	CodeE004: Rejected,
	CodeE005: Rejected,
	CodeR03:  Rejected,
	CodeR12:  Rejected,
	CodeR13:  Rejected,
	CodeR14:  RejectedInvalidCardNumber,
	CodeR16:  Rejected,
	CodeR17:  AmountRateLimitExceeded,
	CodeR58:  Rejected,
}

var StatusByProviderAuthorization = map[string]transactions_factory.OperationStatus{
	// visa - debvisa card present
	Code01 + Visa + CardPresent: RejectedCallForAuthorization, // Refer to card issuer
	Code02 + Visa + CardPresent: RejectedCallForAuthorization,
	Code03 + Visa + CardPresent: Rejected,            // COMERCIANTE INVÁLIDO
	Code04 + Visa + CardPresent: RejectedBlackListed, // RECOLHER CARTÃO (NÃO HÁ FRAUDE)
	Code05 + Visa + CardPresent: Rejected,            // GENERICO
	Code06 + Visa + CardPresent: RejectedFormError,
	Code07 + Visa + CardPresent: RejectedBlackListed, // FRAUDE CONFIRMADA
	Code12 + Visa + CardPresent: Rejected,            // ERRO DE FORMATO (MENSAGERIA)
	Code13 + Visa + CardPresent: Rejected,            // VALOR DA TRANSAÇÃO INVÁLIDA
	Code14 + Visa + CardPresent: RejectedFormError,   // Nº CARTÃO Ñ PERTENCE AO EMISSOR | Nº CARTÃO INVÁLIDO | VIOLAÇÃO DE SEGURANÇA
	Code15 + Visa + CardPresent: Rejected,            // EMISSOR Ñ LOCALIZADO - BIN INCORRETO
	Code19 + Visa + CardPresent: Rejected,
	Code21 + Visa + CardPresent: Rejected,
	Code25 + Visa + CardPresent: Rejected,
	Code28 + Visa + CardPresent: Rejected,
	Code39 + Visa + CardPresent: RejectedFormError,
	Code41 + Visa + CardPresent: RejectedBlackListed,       // CARTÃO PERDIDO
	Code43 + Visa + CardPresent: RejectedBlackListed,       // CARTÃO ROUBADO
	Code46 + Visa + CardPresent: Rejected,                  // Card terminated
	Code51 + Visa + CardPresent: RejectedInsufficientFunds, // SALDO/LIMITE INSUFICIENTE
	Code52 + Visa + CardPresent: RejectedFormError,
	Code53 + Visa + CardPresent: RejectedFormError,
	Code54 + Visa + CardPresent: RejectedInvalidExpirationDate, // CARTÃO VENCIDO / DT EXPIRAÇÃO INVÁLIDA
	Code55 + Visa + CardPresent: RejectedInvalidSecurityCode,   // SENHA INVÁLIDA
	Code57 + Visa + CardPresent: RejectedFormError,             // TRANSAÇÃO NÃO PERMITIDA PARA O CARTÃO
	Code58 + Visa + CardPresent: RejectedCallForAuthorization,  // TRANSAÇÃO NÃO PERMITIDA |CAPACIDADE DO TERMINAL
	Code59 + Visa + CardPresent: RejectedCallForAuthorization,  // SUSPEITA DE FRAUDE
	Code61 + Visa + CardPresent: RejectedInsufficientFunds,     // VALOR EXCESSO | SAQUE
	Code62 + Visa + CardPresent: RejectedCallForAuthorization,  // CARTÃO DOMÉSTICO - TRANSAÇÃO INTERNATIONAL
	Code64 + Visa + CardPresent: RejectedBlackListed,
	Code65 + Visa + CardPresent: AmountRateLimitExceeded, // QUANT. DE SAQUES EXCEDIDO
	Code6P + Visa + CardPresent: RejectedFormError,       // Verification Data Failed
	Code70 + Visa + CardPresent: Rejected,                // PIN data required
	Code74 + Visa + CardPresent: RejectedFormError,
	Code75 + Visa + CardPresent: RejectedMaxAttempts,          // EXCEDIDAS TENTATIVAS DE SENHA | SAQUE
	Code76 + Visa + CardPresent: Rejected,                     // REVERSÃO INVÁLIDA
	Code78 + Visa + CardPresent: DisabledAccount,              // CARTÃO NOVO SEM DESBLOQUEIO
	Code80 + Visa + CardPresent: Rejected,                     // No financial impact (used in reversal responses to declined originals)
	Code81 + Visa + CardPresent: RejectedInvalidSecurityCode,  // SENHA VENCIDA / ERRO DE CRIPTOGRAFIA DE SENHA
	Code82 + Visa + CardPresent: RejectedCallForAuthorization, // CARTÃO INVÁLIDO (criptograma)
	Code85 + Visa + CardPresent: Rejected,                     // No reason to decline a request for address verification, CVV2 verification, or credit voucher or merchandise return
	Code86 + Visa + CardPresent: RejectedInvalidSecurityCode,  // SENHA INVÁLIDA
	Code91 + Visa + CardPresent: Rejected,                     // EMISSOR FORA DO AR
	Code92 + Visa + CardPresent: Rejected,                     // NÃO LOCALIZADO PELO ROTEADOR
	Code93 + Visa + CardPresent: Rejected,
	Code94 + Visa + CardPresent: RejectedDuplicatedPayment, // VALOR DO TRACING DATA DUPLICADO
	Code96 + Visa + CardPresent: Rejected,                  // FALHA DO SISTEMA
	Code1A + Visa + CardPresent: Rejected,
	CodeN0 + Visa + CardPresent: Contingency,
	CodeN3 + Visa + CardPresent: Rejected,
	CodeN4 + Visa + CardPresent: Rejected,
	CodeN7 + Visa + CardPresent: RejectedInvalidSecurityCode, // CVV2 INVALIDO
	CodeN8 + Visa + CardPresent: Rejected,
	CodeQ1 + Visa + CardPresent: Rejected,
	CodeR0 + Visa + CardPresent: Rejected,
	CodeR1 + Visa + CardPresent: Rejected,
	CodeR2 + Visa + CardPresent: Rejected,
	CodeR3 + Visa + CardPresent: Rejected,
	CodeZ1 + Visa + CardPresent: Rejected, // Offline Declined
	CodeZ3 + Visa + CardPresent: Rejected, // Unable to go online; offline-declined

	// visa - debvisa ecommerce
	Code01 + Visa + Ecommerce: RejectedCallForAuthorization, // Refer to card issuer
	Code02 + Visa + Ecommerce: RejectedCallForAuthorization,
	Code03 + Visa + Ecommerce: Rejected,            //  COMERCIANTE INVÁLIDO
	Code04 + Visa + Ecommerce: RejectedBlackListed, //  RECOLHER CARTÃO (NÃO HÁ FRAUDE)
	Code05 + Visa + Ecommerce: Rejected,            //  GENERICO
	Code06 + Visa + Ecommerce: RejectedFormError,
	Code07 + Visa + Ecommerce: RejectedBlackListed,       //  FRAUDE CONFIRMADA
	Code12 + Visa + Ecommerce: Rejected,                  // ERRO DE FORMATO (MENSAGERIA)
	Code13 + Visa + Ecommerce: RejectedFormError,         // VALOR DA TRANSAÇÃO INVÁLIDA
	Code14 + Visa + Ecommerce: RejectedInvalidCardNumber, // Nº CARTÃO Ñ PERTENCE AO EMISSOR | Nº CARTÃO INVÁLIDO | VIOLAÇÃO DE SEGURANÇA
	Code15 + Visa + Ecommerce: Rejected,                  // EMISSOR Ñ LOCALIZADO - BIN INCORRETO
	Code19 + Visa + Ecommerce: Rejected,
	Code21 + Visa + Ecommerce: Rejected,
	Code25 + Visa + Ecommerce: Rejected,
	Code28 + Visa + Ecommerce: Rejected,
	Code39 + Visa + Ecommerce: RejectedFormError,
	Code41 + Visa + Ecommerce: RejectedBlackListed,       // CARTÃO PERDIDO
	Code43 + Visa + Ecommerce: RejectedBlackListed,       // CARTÃO ROUBADO
	Code46 + Visa + Ecommerce: Rejected,                  // Card terminated
	Code51 + Visa + Ecommerce: RejectedInsufficientFunds, // SALDO/LIMITE INSUFICIENTE
	Code52 + Visa + Ecommerce: RejectedFormError,
	Code53 + Visa + Ecommerce: RejectedFormError,
	Code54 + Visa + Ecommerce: RejectedInvalidExpirationDate, // CARTÃO VENCIDO / DT EXPIRAÇÃO INVÁLIDA
	Code55 + Visa + Ecommerce: RejectedInvalidSecurityCode,   // SENHA INVÁLIDA
	Code57 + Visa + Ecommerce: Rejected,                      // TRANSAÇÃO NÃO PERMITIDA PARA O CARTÃO
	Code58 + Visa + Ecommerce: RejectedCallForAuthorization,  // TRANSAÇÃO NÃO PERMITIDA |CAPACIDADE DO TERMINAL
	Code59 + Visa + Ecommerce: RejectedCallForAuthorization,  // SUSPEITA DE FRAUDE
	Code61 + Visa + Ecommerce: RejectedInsufficientFunds,     // VALOR EXCESSO | SAQUE
	Code62 + Visa + Ecommerce: RejectedBlackListed,           // CARTÃO DOMÉSTICO - TRANSAÇÃO INTERNATIONAL
	Code64 + Visa + Ecommerce: RejectedBlackListed,
	Code65 + Visa + Ecommerce: AmountRateLimitExceeded, // QUANT. DE SAQUES EXCEDIDO
	Code6P + Visa + Ecommerce: RejectedFormError,       // Verification Data Failed
	Code70 + Visa + Ecommerce: Rejected,                // PIN data required
	Code74 + Visa + Ecommerce: RejectedFormError,
	Code75 + Visa + Ecommerce: RejectedMaxAttempts,          // EXCEDIDAS TENTATIVAS DE SENHA | SAQUE
	Code76 + Visa + Ecommerce: Rejected,                     // REVERSÃO INVÁLIDA
	Code78 + Visa + Ecommerce: DisabledAccount,              // CARTÃO NOVO SEM DESBLOQUEIO
	Code80 + Visa + Ecommerce: Rejected,                     // No financial impact (used in reversal responses to declined originals)
	Code81 + Visa + Ecommerce: RejectedInvalidSecurityCode,  // SENHA VENCIDA / ERRO DE CRIPTOGRAFIA DE SENHA
	Code82 + Visa + Ecommerce: RejectedCallForAuthorization, // CARTÃO INVÁLIDO (criptograma)
	Code85 + Visa + Ecommerce: Rejected,                     // No reason to decline a request for address verification, CVV2 verification, or credit voucher or merchandise return
	Code86 + Visa + Ecommerce: RejectedInvalidSecurityCode,  // SENHA INVÁLIDA
	Code91 + Visa + Ecommerce: Rejected,                     // EMISSOR FORA DO AR
	Code92 + Visa + Ecommerce: Rejected,                     // NÃO LOCALIZADO PELO ROTEADOR
	Code93 + Visa + Ecommerce: Rejected,
	Code94 + Visa + Ecommerce: RejectedDuplicatedPayment, // VALOR DO TRACING DATA DUPLICADO
	Code96 + Visa + Ecommerce: Rejected,                  // FALHA DO SISTEMA
	Code1A + Visa + Ecommerce: Rejected,
	CodeN0 + Visa + Ecommerce: Contingency,
	CodeN3 + Visa + Ecommerce: Rejected,
	CodeN4 + Visa + Ecommerce: Rejected,
	CodeN7 + Visa + Ecommerce: RejectedInvalidSecurityCode, // CVV2 INVALIDO
	CodeN8 + Visa + Ecommerce: Rejected,
	CodeQ1 + Visa + Ecommerce: Rejected,
	CodeR0 + Visa + Ecommerce: Rejected,
	CodeR1 + Visa + Ecommerce: Rejected,
	CodeR2 + Visa + Ecommerce: Rejected,
	CodeR3 + Visa + Ecommerce: Rejected,
	CodeZ1 + Visa + Ecommerce: Rejected, // Offline Declined
	CodeZ3 + Visa + Ecommerce: Rejected, //  Unable to go online; offline-declined
}

var StatusByProviderAuthorizationCardPresent = map[Code]transactions_factory.OperationStatus{
	Code00:   Approved,
	Code01:   RejectedCallForAuthorization,
	Code02:   RejectedCallForAuthorization,
	Code03:   Rejected,
	Code04:   RejectedBlackListed,
	Code05:   Rejected,
	Code06:   RejectedFormError,
	Code07:   RejectedBlackListed,
	Code08:   Approved,
	Code10:   Approved,
	Code12:   Rejected,
	Code13:   Rejected,
	Code14:   RejectedFormError,
	Code15:   Rejected,
	Code19:   Rejected,
	Code21:   Rejected,
	Code25:   Rejected,
	Code28:   Rejected,
	Code30:   Rejected,
	Code39:   RejectedFormError,
	Code41:   RejectedBlackListed,
	Code43:   RejectedBlackListed,
	Code46:   Rejected,
	Code51:   RejectedInsufficientFunds,
	Code52:   RejectedFormError,
	Code53:   RejectedFormError,
	Code54:   RejectedInvalidExpirationDate,
	Code55:   RejectedInvalidSecurityCode,
	Code57:   RejectedFormError,
	Code58:   RejectedCallForAuthorization,
	Code59:   RejectedBlackListed,
	Code61:   RejectedInsufficientFunds,
	Code62:   RejectedCallForAuthorization,
	Code63:   RejectedCallForAuthorization,
	Code64:   RejectedBlackListed,
	Code65:   AmountRateLimitExceeded,
	Code6P:   RejectedFormError,
	Code70:   RejectedCallForAuthorization,
	Code71:   Rejected,
	Code74:   RejectedFormError,
	Code75:   RejectedMaxAttempts,
	Code76:   Rejected,
	Code77:   Rejected,
	Code78:   DisabledAccount,
	Code79:   RejectedCallForAuthorization,
	Code81:   RejectedInvalidSecurityCode,
	Code82:   RejectedCallForAuthorization,
	Code83:   RejectedCallForAuthorization,
	Code84:   Rejected,
	Code85:   Approved,
	Code86:   RejectedInvalidSecurityCode,
	Code87:   Approved,
	Code88:   RejectedInvalidSecurityCode,
	Code89:   Rejected,
	Code91:   Rejected,
	Code92:   Rejected,
	Code93:   Rejected,
	Code94:   RejectedDuplicatedPayment,
	Code96:   Rejected,
	Code1A:   Rejected,
	CodeN0:   Contingency,
	CodeN3:   Rejected,
	CodeN4:   Rejected,
	CodeN7:   RejectedInvalidSecurityCode,
	CodeN8:   Rejected,
	CodeQ1:   Rejected,
	CodeR0:   Rejected,
	CodeR1:   Rejected,
	CodeR2:   Rejected,
	CodeR3:   Rejected,
	CodeZ1:   Rejected,
	CodeZ3:   Rejected,
	CodeE001: Rejected,
	CodeE002: Contingency,
	CodeE003: RejectedInvalidCardNumber,
	CodeE004: Rejected,
	CodeE005: Rejected,
	CodeR03:  Rejected,
	CodeR12:  Rejected,
	CodeR13:  Rejected,
	CodeR14:  RejectedInvalidCardNumber,
	CodeR58:  Rejected,
}

var StatusByProviderAuthorizationEcommerce = map[Code]transactions_factory.OperationStatus{
	Code00:   Approved,
	Code01:   RejectedCallForAuthorization,
	Code02:   RejectedCallForAuthorization,
	Code03:   Rejected,
	Code04:   RejectedBlackListed,
	Code05:   Rejected,
	Code06:   RejectedFormError,
	Code07:   RejectedBlackListed,
	Code08:   Approved,
	Code10:   Approved,
	Code12:   Rejected,
	Code13:   RejectedFormError,
	Code14:   RejectedInvalidCardNumber,
	Code15:   Rejected,
	Code19:   Rejected,
	Code21:   Rejected,
	Code25:   Rejected,
	Code28:   Rejected,
	Code30:   Rejected,
	Code39:   RejectedFormError,
	Code41:   RejectedBlackListed,
	Code43:   RejectedBlackListed,
	Code46:   Rejected,
	Code51:   RejectedInsufficientFunds,
	Code52:   RejectedFormError,
	Code53:   RejectedFormError,
	Code54:   RejectedInvalidExpirationDate,
	Code55:   RejectedInvalidSecurityCode,
	Code57:   Rejected,
	Code58:   RejectedCallForAuthorization,
	Code59:   RejectedBlackListed,
	Code61:   RejectedInsufficientFunds,
	Code62:   RejectedBlackListed,
	Code63:   RejectedCallForAuthorization,
	Code64:   RejectedBlackListed,
	Code65:   AmountRateLimitExceeded,
	Code6P:   RejectedFormError,
	Code70:   RejectedCallForAuthorization,
	Code71:   Rejected,
	Code74:   RejectedFormError,
	Code75:   RejectedMaxAttempts,
	Code76:   Rejected,
	Code77:   Rejected,
	Code78:   DisabledAccount,
	Code79:   RejectedCallForAuthorization,
	Code81:   RejectedInvalidSecurityCode,
	Code82:   RejectedCallForAuthorization,
	Code83:   RejectedCallForAuthorization,
	Code84:   Rejected,
	Code85:   Approved,
	Code86:   RejectedInvalidSecurityCode,
	Code87:   Approved,
	Code88:   RejectedInvalidSecurityCode,
	Code89:   Rejected,
	Code91:   Rejected,
	Code92:   Rejected,
	Code93:   Rejected,
	Code94:   RejectedCallForAuthorization,
	Code96:   Rejected,
	Code1A:   Rejected,
	CodeN0:   Contingency,
	CodeN3:   Rejected,
	CodeN4:   Rejected,
	CodeN7:   RejectedInvalidSecurityCode,
	CodeN8:   Rejected,
	CodeQ1:   Rejected,
	CodeR0:   Rejected,
	CodeR1:   Rejected,
	CodeR2:   Rejected,
	CodeR3:   Rejected,
	CodeZ1:   Rejected,
	CodeZ3:   Rejected,
	CodeE001: Rejected,
	CodeE002: Contingency,
	CodeE003: RejectedInvalidCardNumber,
	CodeE004: Rejected,
	CodeE005: Rejected,
	CodeR03:  Rejected,
	CodeR12:  Rejected,
	CodeR13:  Rejected,
	CodeR14:  RejectedInvalidCardNumber,
	CodeR58:  Rejected,
}

var StatusByProviderPurchase = map[Code]transactions_factory.OperationStatus{
	Code00:   Approved,
	Code01:   RejectedCallForAuthorization,
	Code02:   RejectedCallForAuthorization,
	Code03:   Rejected,
	Code04:   RejectedBlackListed,
	Code05:   Rejected,
	Code06:   RejectedFormError,
	Code07:   RejectedBlackListed,
	Code08:   Approved,
	Code10:   Approved,
	Code12:   Rejected,
	Code13:   RejectedFormError,
	Code14:   RejectedInvalidCardNumber,
	Code15:   Rejected,
	Code19:   Rejected,
	Code21:   Rejected,
	Code25:   Rejected,
	Code28:   Rejected,
	Code30:   Rejected,
	Code39:   RejectedFormError,
	Code41:   RejectedBlackListed,
	Code43:   RejectedBlackListed,
	Code46:   Rejected,
	Code51:   RejectedInsufficientFunds,
	Code52:   RejectedFormError,
	Code53:   RejectedFormError,
	Code54:   RejectedInvalidExpirationDate,
	Code55:   RejectedInvalidSecurityCode,
	Code57:   Rejected,
	Code58:   RejectedCallForAuthorization,
	Code59:   RejectedBlackListed,
	Code61:   RejectedInsufficientFunds,
	Code62:   RejectedBlackListed,
	Code63:   RejectedCallForAuthorization,
	Code64:   RejectedBlackListed,
	Code65:   AmountRateLimitExceeded,
	Code6P:   RejectedFormError,
	Code70:   RejectedCallForAuthorization,
	Code71:   Rejected,
	Code74:   RejectedFormError,
	Code75:   RejectedMaxAttempts,
	Code76:   Rejected,
	Code77:   Rejected,
	Code78:   DisabledAccount,
	Code79:   RejectedCallForAuthorization,
	Code81:   RejectedInvalidSecurityCode,
	Code82:   RejectedCallForAuthorization,
	Code83:   RejectedCallForAuthorization,
	Code84:   Rejected,
	Code85:   Approved,
	Code86:   RejectedInvalidSecurityCode,
	Code87:   Approved,
	Code88:   RejectedInvalidSecurityCode,
	Code89:   Rejected,
	Code91:   Rejected,
	Code92:   Rejected,
	Code93:   Rejected,
	Code94:   RejectedCallForAuthorization,
	Code96:   Rejected,
	Code1A:   Rejected,
	CodeN0:   Contingency,
	CodeN3:   Rejected,
	CodeN4:   Rejected,
	CodeN7:   RejectedInvalidSecurityCode,
	CodeN8:   Rejected,
	CodeQ1:   Rejected,
	CodeR0:   Rejected,
	CodeR1:   Rejected,
	CodeR2:   Rejected,
	CodeR3:   Rejected,
	CodeZ1:   Rejected,
	CodeZ3:   Rejected,
	CodeE001: Rejected,
	CodeE002: Contingency,
	CodeE003: RejectedInvalidCardNumber,
	CodeE004: Rejected,
	CodeE005: Rejected,
	CodeR03:  Rejected,
	CodeR12:  Rejected,
	CodeR13:  Rejected,
	CodeR14:  RejectedInvalidCardNumber,
	CodeR58:  Rejected,
}

type Code = string

//  Response codes are used to indicate approval or decline of a transaction.
const (
	Code00 Code = "00" // Approved or completed successfully
	Code01 Code = "01" // Refer to card issuer
	Code02 Code = "02" // Refer to card issuer, special condition
	Code03 Code = "03" // Invalid merchant
	Code04 Code = "04" // Capture Card
	Code05 Code = "05" // Do not honor
	Code06 Code = "06" // Error / Format error
	Code07 Code = "07" // Pick up card, special condition (fraud account)
	Code08 Code = "08" // Honor with ID
	Code10 Code = "10" // Partial Approval
	Code12 Code = "12" // Invalid transaction
	Code13 Code = "13" // Invalid amount
	Code14 Code = "14" // Invalid card number
	Code15 Code = "15" // Invalid issuer
	Code19 Code = "19" // Problem in Acquirer
	Code21 Code = "21" // No action taken
	Code25 Code = "25" // Unable to locate record in file
	Code28 Code = "28" // File is temporarily unavailable for update or inquiry
	Code30 Code = "30" // Format error
	Code39 Code = "39" // No credit account
	Code41 Code = "41" // Lost card
	Code43 Code = "43" // Stolen card
	Code46 Code = "46" // Closed Account
	Code51 Code = "51" // Insufficient funds/over credit limit
	Code52 Code = "52" // No checking account
	Code53 Code = "53" // No savings account
	Code54 Code = "54" // Expired card
	Code55 Code = "55" // Invalid PIN
	Code57 Code = "57" // Transaction not permitted to issuer/cardholder
	Code59 Code = "59" // Suspected fraud
	Code58 Code = "58" // Transaction not permitted to acquirer/terminal
	Code61 Code = "61" // Exceeds withdrawal amount limit
	Code62 Code = "62" // Restricted card
	Code63 Code = "63" // Security violation
	Code64 Code = "64" // Transaction does not fulfill AML requirement
	Code65 Code = "65" // Exceeds withdrawal count limit
	Code70 Code = "70" // Contact Card Issuer
	Code71 Code = "71" // PIN Not Changed
	Code74 Code = "74" // Different value than that used for PIN encryption errors
	Code75 Code = "75" // Allowable number of PIN tries exceeded
	Code76 Code = "76" // Invalid/nonexistent “To Account” specified
	Code77 Code = "77" // Invalid/nonexistent “From Account” specified
	Code78 Code = "78" // Invalid/nonexistent account specified (general)
	Code79 Code = "79" // Lifecycle (Mastercard use only)
	Code80 Code = "80" // System not available
	Code81 Code = "81" // Domestic Debit Transaction Not Allowed (Regional useonly)
	Code82 Code = "82" // Visa: Negative Online or offline PIN authentication interrupted, Mastercard: Policy (Mastercard use only)
	Code83 Code = "83" // Security (Mastercard use only)
	Code84 Code = "84" // Invalid Authorization Life Cycle
	Code85 Code = "85" // Not declined Valid for all zero amount transactions.
	Code86 Code = "86" // PIN Validation not possible
	Code87 Code = "87" // Purchase Amount Only, No Cash Back Allowed
	Code88 Code = "88" // Cryptographic failure
	Code89 Code = "89" // Unacceptable PIN— Transaction Declined—Retry
	Code91 Code = "91" // Authorization System or issuer system inoperative
	Code92 Code = "92" // Unable to route transaction
	Code93 Code = "93" // Transaction cannot be completed - violation of law.
	Code94 Code = "94" // Duplicate transmission detected
	Code96 Code = "96" // Resposta de erro do sistema indica que o emissor do cartão do cliente está tendo dificuldades para processar o pagamento
	Code1A Code = "1A" // Additional customer authentication required
	CodeN0 Code = "N0" // Force STIP
	CodeN3 Code = "N3" // Cash service not available
	CodeN4 Code = "N4" // Cash request exceeds issuer or Approved limit
	CodeN7 Code = "N7" // Decline for CVV2 failure
	CodeN8 Code = "N8" // Transaction amount exceeds pre-authorized approval amount
	CodeP5 Code = "P5" // Denied PIN unblock-PIN change or unblock request declined by issuer
	CodeP6 Code = "P6" // Denied PIN change-requested PIN unsafe
	CodeQ1 Code = "Q1" // Card authentication failed / Offline PIN authentication interrupted
	CodeR0 Code = "R0" // Stop payment order
	CodeR1 Code = "R1" // Revocation of authorization order
	CodeR2 Code = "R2" // Transaction does not qualify for Visa PIN
	CodeR3 Code = "R3" // Revocation of all authorizations order
	Code6P Code = "6P" // Verification Data Failed
	CodeZ1 Code = "Z1" // Offline Declined
	CodeZ3 Code = "Z3" // Unable to go online

	CodeE001 Code = "E001" // Rejected
	CodeE002 Code = "E002" // Rejected - System error
	CodeE003 Code = "E003" // Invalid card
	CodeE004 Code = "E004" // "Rejected - Invalid operation transaction"
	CodeE005 Code = "E005" // Invalid merchant

	//  Acquirer rejection codes
	CodeR03 Code = "R03" // Invalid Merchant
	CodeR12 Code = "R12" // Invalid transaction
	CodeR13 Code = "R13" // Invalid amount
	CodeR14 Code = "R14" // Invalid card number
	CodeR16 Code = "R16" // Cancel requests limit reached
	CodeR17 Code = "R17" // Refund amount exceed the transaction amount
	CodeR58 Code = "R58" // Transaction not permitted to acquirer/terminal
)

const (
	Visa        = "_visa"
	Master      = "_master"
	CardPresent = "_card_present"
	Ecommerce   = "_ecommerce"
)

var LimitRetryNumber = map[string]uint32{
	"visa":    15,
	"debvisa": 15,
}
