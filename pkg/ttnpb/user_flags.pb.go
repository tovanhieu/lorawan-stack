// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v1.0.2
// - protoc              v3.9.1
// source: lorawan-stack/api/user.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	gogo "github.com/TheThingsIndustries/protoc-gen-go-flags/gogo"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForUser adds flags to select fields in User.
func AddSelectFlagsForUser(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("deleted-at", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("name", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("name", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("description", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("description", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("attributes", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("attributes", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("contact-info", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("contact-info", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("primary-email-address", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("primary-email-address", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("primary-email-address-validated-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("primary-email-address-validated-at", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("password", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("password", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("password-updated-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("password-updated-at", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("require-password-update", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("require-password-update", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("state", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("state", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("state-description", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("state-description", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("admin", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("admin", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("temporary-password", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("temporary-password", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("temporary-password-created-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("temporary-password-created-at", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("temporary-password-expires-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("temporary-password-expires-at", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("profile-picture", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("profile-picture", prefix), true), flagsplugin.WithHidden(hidden)))
	// NOTE: profile_picture (Picture) does not seem to have select flags.
}

// SelectFromFlags outputs the fieldmask paths forUser message from select flags.
func PathsFromSelectFlagsForUser(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("deleted_at", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("name", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("name", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("description", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("description", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("attributes", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("attributes", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("contact_info", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("contact_info", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("primary_email_address", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("primary_email_address", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("primary_email_address_validated_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("primary_email_address_validated_at", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("password", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("password", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("password_updated_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("password_updated_at", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("require_password_update", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("require_password_update", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("state", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("state", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("state_description", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("state_description", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("admin", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("admin", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("temporary_password", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("temporary_password", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("temporary_password_created_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("temporary_password_created_at", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("temporary_password_expires_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("temporary_password_expires_at", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("profile_picture", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("profile_picture", prefix))
	}
	// NOTE: profile_picture (Picture) does not seem to have select flags.
	return paths, nil
}

// AddSetFlagsForUser adds flags to select fields in User.
func AddSetFlagsForUser(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForUserIdentifiers(flags, flagsplugin.Prefix("ids", prefix), true)
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes", prefix), "", flagsplugin.WithHidden(hidden)))
	// FIXME: Skipping ContactInfo because repeated messages are currently not supported.
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("primary-email-address", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("primary-email-address-validated-at", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("password", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("require-password-update", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("state", prefix), flagsplugin.EnumValueDesc(State_value), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("state-description", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("admin", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("temporary-password", prefix), "", flagsplugin.WithHidden(hidden)))
	// FIXME: Skipping ProfilePicture because it does not seem to implement AddSetFlags.
}

// SetFromFlags sets the User message from flags.
func (m *User) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("ids", prefix)); changed {
		if m.Ids == nil {
			m.Ids = &UserIdentifiers{}
		}
		if setPaths, err := m.Ids.SetFromFlags(flags, flagsplugin.Prefix("ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Name = val
		paths = append(paths, flagsplugin.Prefix("name", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Description = val
		paths = append(paths, flagsplugin.Prefix("description", prefix))
	}
	if val, changed, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Attributes = val
		paths = append(paths, flagsplugin.Prefix("attributes", prefix))
	}
	// FIXME: Skipping ContactInfo because it does not seem to implement AddSetFlags.
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("primary_email_address", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.PrimaryEmailAddress = val
		paths = append(paths, flagsplugin.Prefix("primary_email_address", prefix))
	}
	if val, changed, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("primary_email_address_validated_at", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.PrimaryEmailAddressValidatedAt = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("primary_email_address_validated_at", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("password", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Password = val
		paths = append(paths, flagsplugin.Prefix("password", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("require_password_update", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.RequirePasswordUpdate = val
		paths = append(paths, flagsplugin.Prefix("require_password_update", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("state", prefix)); err != nil {
		return nil, err
	} else if changed {
		enumValue, err := flagsplugin.SetEnumString(val, State_value)
		if err != nil {
			return nil, err
		}
		m.State = State(enumValue)
		paths = append(paths, flagsplugin.Prefix("state", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("state_description", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.StateDescription = val
		paths = append(paths, flagsplugin.Prefix("state_description", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("admin", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Admin = val
		paths = append(paths, flagsplugin.Prefix("admin", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("temporary_password", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.TemporaryPassword = val
		paths = append(paths, flagsplugin.Prefix("temporary_password", prefix))
	}
	// FIXME: Skipping ProfilePicture because it does not seem to implement AddSetFlags.
	return paths, nil
}

// AddSelectFlagsForListUsersRequest adds flags to select fields in ListUsersRequest.
func AddSelectFlagsForListUsersRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("field-mask", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("field-mask", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("order", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("order", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("limit", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("limit", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("page", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("page", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("deleted", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forListUsersRequest message from select flags.
func PathsFromSelectFlagsForListUsersRequest(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}

// AddSetFlagsForListUsersRequest adds flags to select fields in ListUsersRequest.
func AddSetFlagsForListUsersRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the ListUsersRequest message from flags.
func (m *ListUsersRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Deleted = val
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}
