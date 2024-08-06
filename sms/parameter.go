package sendcloud

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"time"
)

func (client *SendCloudSms) calculateSignature(params url.Values) string {
	sortedParams := url.Values{}

	for k, v := range params {
		if k != "smsKey" && k != "signature" {
			sortedParams[k] = v
		}
	}

	keys := make([]string, 0, len(sortedParams))
	for k := range sortedParams {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var paramStr string
	for _, k := range keys {
		paramStr += k + "=" + sortedParams.Get(k) + "&"
	}

	if len(paramStr) > 0 {
		paramStr = paramStr[:len(paramStr)-1]
	}

	signStr := client.smsKey + "&" + paramStr + "&" + client.smsKey

	hasher := sha256.New()
	hasher.Write([]byte(signStr))
	sha256Bytes := hasher.Sum(nil)

	signature := hex.EncodeToString(sha256Bytes)

	return signature
}

func (client *SendCloudSms) prepareSendTemplateSmsParams(args *TemplateSms) (url.Values, error) {
	params := url.Values{}
	params.Set("smsUser", client.smsUser)
	params.Set("msgType", strconv.Itoa(args.MsgType))
	params.Set("phone", args.Phone)
	params.Set("templateId", strconv.Itoa(args.TemplateId))
	params.Set("timestamp", strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10))
	if len(args.Vars) > 0 {
		varsJSON, err := json.Marshal(args.Vars)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal vars: %v", err)
		}
		params.Set("vars", string(varsJSON))
	}
	if args.LabelId != 0 {
		params.Set("labelId", strconv.Itoa(args.LabelId))
	}
	if len(args.SendRequestId) > 0 {
		params.Set("sendRequestId", args.SendRequestId)
	}
	if len(args.Tag) > 0 {
		tagJSON, err := json.Marshal(args.Tag)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal tag: %v", err)
		}
		params.Set("tag", string(tagJSON))
	}
	return params, nil
}

func (client *SendCloudSms) prepareSendVoiceSmsParams(args *VoiceSms) (url.Values, error) {
	params := url.Values{}
	params.Set("smsUser", client.smsUser)
	params.Set("phone", args.Phone)
	params.Set("code", args.Code)
	if args.LabelId != 0 {
		params.Set("labelId", strconv.Itoa(args.LabelId))
	}
	if len(args.SendRequestId) > 0 {
		params.Set("sendRequestId", args.SendRequestId)
	}
	params.Set("timestamp", strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10))
	if len(args.Tag) > 0 {
		tagJSON, err := json.Marshal(args.Tag)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal tag: %v", err)
		}
		params.Set("tag", string(tagJSON))
	}
	return params, nil
}

func (client *SendCloudSms) prepareSendCodeSmsParams(args *CodeSms) (url.Values, error) {
	params := url.Values{}
	params.Set("smsUser", client.smsUser)
	params.Set("msgType", strconv.Itoa(args.MsgType))
	params.Set("phone", args.Phone)
	if args.SignId != 0 {
		params.Set("signId", strconv.Itoa(args.SignId))
	}
	if len(args.SignName) > 0 {
		params.Set("signName", args.SignName)
	}
	params.Set("code", args.Code)
	if args.LabelId != 0 {
		params.Set("labelId", strconv.Itoa(args.LabelId))
	}
	if len(args.SendRequestId) > 0 {
		params.Set("sendRequestId", args.SendRequestId)
	}
	params.Set("timestamp", strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10))
	if len(args.Tag) > 0 {
		tagJSON, err := json.Marshal(args.Tag)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal tag: %v", err)
		}
		params.Set("tag", string(tagJSON))
	}
	return params, nil
}
