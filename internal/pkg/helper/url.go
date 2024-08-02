package helper

import "fmt"

func NFTAssetURL(baseURL string, collectionAddress string, tokenID string) string {
	return fmt.Sprintf("%s/asset/%s/%s", baseURL, collectionAddress, tokenID)
}

func AuctionBidURL(baseURL string, collectionAddress string, tokenID string) string {
	return fmt.Sprintf("%s/auctions/bid/%s/%s", baseURL, collectionAddress, tokenID)
}

func DepositURL(baseURL string) string {
	return fmt.Sprintf("%s/liquidity/deposit-eth", baseURL)
}

func MyNFTsURL(baseURL string) string {
	return fmt.Sprintf("%s/account/my-nfts", baseURL)
}

func MyBorrowsURL(baseURL string) string {
	return fmt.Sprintf("%s/account/my-nfts?type_in=in_loan", baseURL)
}

func BatchBorrowURL(baseURL string) string {
	return fmt.Sprintf("%s/account/my-nfts/?action=batch-borrow", baseURL)
}

func WithdrawURL(baseURL string, reserveAddress string, lendPoolAddressesProviderAddress string) string {
	return fmt.Sprintf("%s/withdraw/%s%s", baseURL, reserveAddress, lendPoolAddressesProviderAddress)
}

func RepayURL(baseURL string, collectionAddress string, tokenID string) string {
	return fmt.Sprintf("%s/liquidity/batch-repay", baseURL)
}

func AuctionHistoryURL(baseURL string) string {
	return fmt.Sprintf("%s/auctions/auction-history", baseURL)
}

func MyActiveAuctionURL(baseURL string) string {
	return fmt.Sprintf("%s/account/auctions", baseURL)
}

func AvailableToAuctionURL(baseURL string) string {
	return fmt.Sprintf("%s/en/auctions/available-to-auction", baseURL)
}

func LoansInAuctionURL(baseURL string) string {
	return fmt.Sprintf("%s/auctions/loans-in-auction", baseURL)
}

func HealthFactorAlertListURL(baseURL string) string {
	return fmt.Sprintf("%s/auctions/health-factor-alert-list", baseURL)
}

func LiquidateAuctionURL(baseURL string, collectionAddress string, tokenID string) string {
	return fmt.Sprintf("%s/auction/bid/%s/%s", baseURL, collectionAddress, tokenID)
}

func EmailConfirmURL(baseURL string, code string) string {
	return fmt.Sprintf("%s/email-confirm/%s", baseURL, code)
}

func MyPairingListingsURL(baseURL string) string {
	return fmt.Sprintf("%s/apecoin-staking/my-pairing-listings/", baseURL)
}

func APEsURL(baseURL string) string {
	return fmt.Sprintf("%s/apecoin-staking/my-apes/", baseURL)
}

func StakingsURL(baseURL string) string {
	return fmt.Sprintf("%s/apecoin-staking/my-stakings/", baseURL)
}

func PairingListingsURL(baseURL string) string {
	return fmt.Sprintf("%s/apecoin-staking/explore/pairing-listings/", baseURL)
}

func ApeCoinImageUrl() string {
	return "https://www.benddao.xyz/images/elements/apecoin-xl.png"
}
