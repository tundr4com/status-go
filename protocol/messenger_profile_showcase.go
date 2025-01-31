package protocol

import (
	"crypto/ecdsa"
	crand "crypto/rand"
	"errors"
	"reflect"

	"google.golang.org/protobuf/proto"

	"github.com/status-im/status-go/protocol/common"
	"github.com/status-im/status-go/protocol/protobuf"
)

func toProfileShowcaseCommunityProto(preferences []*ProfileShowcaseCommunityPreference, visibility ProfileShowcaseVisibility) []*protobuf.ProfileShowcaseCommunity {
	communities := []*protobuf.ProfileShowcaseCommunity{}
	for _, preference := range preferences {
		if preference.ShowcaseVisibility != visibility {
			continue
		}

		communities = append(communities, &protobuf.ProfileShowcaseCommunity{
			CommunityId: preference.CommunityID,
			Order:       uint32(preference.Order),
		})
	}
	return communities
}

func toProfileShowcaseAccountProto(preferences []*ProfileShowcaseAccountPreference, visibility ProfileShowcaseVisibility) []*protobuf.ProfileShowcaseAccount {
	accounts := []*protobuf.ProfileShowcaseAccount{}
	for _, preference := range preferences {
		if preference.ShowcaseVisibility != visibility {
			continue
		}

		accounts = append(accounts, &protobuf.ProfileShowcaseAccount{
			Address: preference.Address,
			Name:    preference.Name,
			ColorId: preference.ColorID,
			Emoji:   preference.Emoji,
			Order:   uint32(preference.Order),
		})
	}
	return accounts
}

func toProfileShowcaseCollectibleProto(preferences []*ProfileShowcaseCollectiblePreference, visibility ProfileShowcaseVisibility) []*protobuf.ProfileShowcaseCollectible {
	collectibles := []*protobuf.ProfileShowcaseCollectible{}
	for _, preference := range preferences {
		if preference.ShowcaseVisibility != visibility {
			continue
		}

		collectibles = append(collectibles, &protobuf.ProfileShowcaseCollectible{
			Uid:   preference.UID,
			Order: uint32(preference.Order),
		})
	}
	return collectibles
}

func toProfileShowcaseAssetProto(preferences []*ProfileShowcaseAssetPreference, visibility ProfileShowcaseVisibility) []*protobuf.ProfileShowcaseAsset {
	assets := []*protobuf.ProfileShowcaseAsset{}
	for _, preference := range preferences {
		if preference.ShowcaseVisibility != visibility {
			continue
		}

		assets = append(assets, &protobuf.ProfileShowcaseAsset{
			Symbol: preference.Symbol,
			Order:  uint32(preference.Order),
		})
	}
	return assets
}

func fromProfileShowcaseCommunityProto(messages []*protobuf.ProfileShowcaseCommunity) []*ProfileShowcaseCommunity {
	communities := []*ProfileShowcaseCommunity{}
	for _, entry := range messages {
		communities = append(communities, &ProfileShowcaseCommunity{
			CommunityID: entry.CommunityId,
			Order:       int(entry.Order),
		})
	}
	return communities
}

func fromProfileShowcaseAccountProto(messages []*protobuf.ProfileShowcaseAccount) []*ProfileShowcaseAccount {
	accounts := []*ProfileShowcaseAccount{}
	for _, entry := range messages {
		accounts = append(accounts, &ProfileShowcaseAccount{
			Address: entry.Address,
			Name:    entry.Name,
			ColorID: entry.ColorId,
			Emoji:   entry.Emoji,
			Order:   int(entry.Order),
		})
	}
	return accounts
}

func fromProfileShowcaseCollectibleProto(messages []*protobuf.ProfileShowcaseCollectible) []*ProfileShowcaseCollectible {
	collectibles := []*ProfileShowcaseCollectible{}
	for _, entry := range messages {
		collectibles = append(collectibles, &ProfileShowcaseCollectible{
			UID:   entry.Uid,
			Order: int(entry.Order),
		})
	}
	return collectibles
}

func fromProfileShowcaseAssetProto(messages []*protobuf.ProfileShowcaseAsset) []*ProfileShowcaseAsset {
	assets := []*ProfileShowcaseAsset{}
	for _, entry := range messages {
		assets = append(assets, &ProfileShowcaseAsset{
			Symbol: entry.Symbol,
			Order:  int(entry.Order),
		})
	}
	return assets
}

func (m *Messenger) SetProfileShowcasePreferences(preferences *ProfileShowcasePreferences) error {
	err := m.persistence.SaveProfileShowcasePreferences(preferences)
	if err != nil {
		return err
	}

	return m.publishContactCode()
}

func (m *Messenger) GetProfileShowcasePreferences() (*ProfileShowcasePreferences, error) {
	return m.persistence.GetProfileShowcasePreferences()
}

func (m *Messenger) GetProfileShowcaseForContact(contactID string) (*ProfileShowcase, error) {
	return m.persistence.GetProfileShowcaseForContact(contactID)
}

func (m *Messenger) EncryptProfileShowcaseEntriesWithContactPubKeys(entries *protobuf.ProfileShowcaseEntries, contacts []*Contact) (*protobuf.ProfileShowcaseEntriesEncrypted, error) {
	// Make AES key
	AESKey := make([]byte, 32)
	_, err := crand.Read(AESKey)
	if err != nil {
		return nil, err
	}

	// Encrypt showcase entries with the AES key
	data, err := proto.Marshal(entries)
	if err != nil {
		return nil, err
	}

	encrypted, err := common.Encrypt(data, AESKey, crand.Reader)
	if err != nil {
		return nil, err
	}

	eAESKeys := [][]byte{}
	// Sign for each contact
	for _, contact := range contacts {
		var pubK *ecdsa.PublicKey
		var sharedKey []byte
		var eAESKey []byte

		pubK, err = contact.PublicKey()
		if err != nil {
			return nil, err
		}
		// Generate a Diffie-Helman (DH) between the sender private key and the recipient's public key
		sharedKey, err = common.MakeECDHSharedKey(m.identity, pubK)
		if err != nil {
			return nil, err
		}

		// Encrypt the main AES key with AES encryption using the DH key
		eAESKey, err = common.Encrypt(AESKey, sharedKey, crand.Reader)
		if err != nil {
			return nil, err
		}

		eAESKeys = append(eAESKeys, eAESKey)
	}

	return &protobuf.ProfileShowcaseEntriesEncrypted{
		EncryptedEntries: encrypted,
		EncryptionKeys:   eAESKeys,
	}, nil
}

func (m *Messenger) DecryptProfileShowcaseEntriesWithPubKey(senderPubKey *ecdsa.PublicKey, encrypted *protobuf.ProfileShowcaseEntriesEncrypted) (*protobuf.ProfileShowcaseEntries, error) {
	for _, eAESKey := range encrypted.EncryptionKeys {
		// Generate a Diffie-Helman (DH) between the recipient's private key and the sender's public key
		sharedKey, err := common.MakeECDHSharedKey(m.identity, senderPubKey)
		if err != nil {
			return nil, err
		}

		// Decrypt the main encryption AES key with AES encryption using the DH key
		dAESKey, err := common.Decrypt(eAESKey, sharedKey)
		if err != nil {
			if err.Error() == ErrCipherMessageAutentificationFailed {
				continue
			}
			return nil, err
		}
		if dAESKey == nil {
			return nil, errors.New("decrypting the payload encryption key resulted in no error and a nil key")
		}

		// Decrypt profile entries with the newly decrypted main encryption AES key
		entriesData, err := common.Decrypt(encrypted.EncryptedEntries, dAESKey)
		if err != nil {
			return nil, err
		}

		entries := &protobuf.ProfileShowcaseEntries{}
		err = proto.Unmarshal(entriesData, entries)
		if err != nil {
			return nil, err
		}

		return entries, nil
	}

	// Return empty if no matching key found
	return &protobuf.ProfileShowcaseEntries{}, nil
}

func (m *Messenger) GetProfileShowcaseForSelfIdentity() (*protobuf.ProfileShowcase, error) {
	preferences, err := m.GetProfileShowcasePreferences()
	if err != nil {
		return nil, err
	}

	forEveryone := &protobuf.ProfileShowcaseEntries{
		Communities:  toProfileShowcaseCommunityProto(preferences.Communities, ProfileShowcaseVisibilityEveryone),
		Accounts:     toProfileShowcaseAccountProto(preferences.Accounts, ProfileShowcaseVisibilityEveryone),
		Collectibles: toProfileShowcaseCollectibleProto(preferences.Collectibles, ProfileShowcaseVisibilityEveryone),
		Assets:       toProfileShowcaseAssetProto(preferences.Assets, ProfileShowcaseVisibilityEveryone),
	}

	forContacts := &protobuf.ProfileShowcaseEntries{
		Communities:  toProfileShowcaseCommunityProto(preferences.Communities, ProfileShowcaseVisibilityContacts),
		Accounts:     toProfileShowcaseAccountProto(preferences.Accounts, ProfileShowcaseVisibilityContacts),
		Collectibles: toProfileShowcaseCollectibleProto(preferences.Collectibles, ProfileShowcaseVisibilityContacts),
		Assets:       toProfileShowcaseAssetProto(preferences.Assets, ProfileShowcaseVisibilityContacts),
	}

	forIDVerifiedContacts := &protobuf.ProfileShowcaseEntries{
		Communities:  toProfileShowcaseCommunityProto(preferences.Communities, ProfileShowcaseVisibilityIDVerifiedContacts),
		Accounts:     toProfileShowcaseAccountProto(preferences.Accounts, ProfileShowcaseVisibilityIDVerifiedContacts),
		Collectibles: toProfileShowcaseCollectibleProto(preferences.Collectibles, ProfileShowcaseVisibilityIDVerifiedContacts),
		Assets:       toProfileShowcaseAssetProto(preferences.Assets, ProfileShowcaseVisibilityIDVerifiedContacts),
	}

	mutualContacts := []*Contact{}
	iDVerifiedContacts := []*Contact{}

	m.allContacts.Range(func(_ string, contact *Contact) (shouldContinue bool) {
		if contact.mutual() {
			mutualContacts = append(mutualContacts, contact)
			if contact.IsVerified() {
				iDVerifiedContacts = append(iDVerifiedContacts, contact)
			}
		}
		return true
	})

	forContactsEncrypted, err := m.EncryptProfileShowcaseEntriesWithContactPubKeys(forContacts, mutualContacts)
	if err != nil {
		return nil, err
	}

	forIDVerifiedContactsEncrypted, err := m.EncryptProfileShowcaseEntriesWithContactPubKeys(forIDVerifiedContacts, iDVerifiedContacts)
	if err != nil {
		return nil, err
	}

	return &protobuf.ProfileShowcase{
		ForEveryone:           forEveryone,
		ForContacts:           forContactsEncrypted,
		ForIdVerifiedContacts: forIDVerifiedContactsEncrypted,
	}, nil
}

func (m *Messenger) BuildProfileShowcaseFromIdentity(state *ReceivedMessageState, message *protobuf.ProfileShowcase) error {
	communities := []*ProfileShowcaseCommunity{}
	accounts := []*ProfileShowcaseAccount{}
	collectibles := []*ProfileShowcaseCollectible{}
	assets := []*ProfileShowcaseAsset{}

	communities = append(communities, fromProfileShowcaseCommunityProto(message.ForEveryone.Communities)...)
	accounts = append(accounts, fromProfileShowcaseAccountProto(message.ForEveryone.Accounts)...)
	collectibles = append(collectibles, fromProfileShowcaseCollectibleProto(message.ForEveryone.Collectibles)...)
	assets = append(assets, fromProfileShowcaseAssetProto(message.ForEveryone.Assets)...)

	senderPubKey := state.CurrentMessageState.PublicKey
	contactID := state.CurrentMessageState.Contact.ID

	forContacts, err := m.DecryptProfileShowcaseEntriesWithPubKey(senderPubKey, message.ForContacts)
	if err != nil {
		return err
	}

	if forContacts != nil {
		communities = append(communities, fromProfileShowcaseCommunityProto(forContacts.Communities)...)
		accounts = append(accounts, fromProfileShowcaseAccountProto(forContacts.Accounts)...)
		collectibles = append(collectibles, fromProfileShowcaseCollectibleProto(forContacts.Collectibles)...)
		assets = append(assets, fromProfileShowcaseAssetProto(forContacts.Assets)...)
	}

	forIDVerifiedContacts, err := m.DecryptProfileShowcaseEntriesWithPubKey(senderPubKey, message.ForIdVerifiedContacts)
	if err != nil {
		return err
	}

	if forIDVerifiedContacts != nil {
		communities = append(communities, fromProfileShowcaseCommunityProto(forIDVerifiedContacts.Communities)...)
		accounts = append(accounts, fromProfileShowcaseAccountProto(forIDVerifiedContacts.Accounts)...)
		collectibles = append(collectibles, fromProfileShowcaseCollectibleProto(forIDVerifiedContacts.Collectibles)...)
		assets = append(assets, fromProfileShowcaseAssetProto(forIDVerifiedContacts.Assets)...)
	}

	newShowcase := &ProfileShowcase{
		ContactID:    contactID,
		Communities:  communities,
		Accounts:     accounts,
		Collectibles: collectibles,
		Assets:       assets,
	}

	oldShowcase, err := m.persistence.GetProfileShowcaseForContact(contactID)
	if err != nil {
		return err
	}

	if reflect.DeepEqual(newShowcase, oldShowcase) {
		return nil
	}

	err = m.persistence.ClearProfileShowcaseForContact(contactID)
	if err != nil {
		return err
	}

	err = m.persistence.SaveProfileShowcaseForContact(newShowcase)
	if err != nil {
		return err
	}

	state.Response.AddProfileShowcase(newShowcase)
	return nil
}
