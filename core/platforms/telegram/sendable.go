package telegram

type sendable interface {
	Send() (*MessageResponse, error)
}

// Send sends the message.
// On success, the sent message is returned as a MessageResponse.
func (om *OutgoingMessage) Send() (*MessageResponse, error) {
	return om.api.send(om)
}

// Send sends the location.
// On success, the sent message is returned as a MessageResponse.
func (ol *OutgoingLocation) Send() (*MessageResponse, error) {
	return ol.api.send(ol)
}

// Send sends the forward.
// On success, the sent message is returned as a MessageResponse.
func (of *OutgoingForward) Send() (*MessageResponse, error) {
	return of.api.send(of)
}

// Send sends the video.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what bots can send, please check the API documentation.
// On success, the sent message is returned as a MessageResponse.
func (ov *OutgoingVideo) Send() (*MessageResponse, error) {
	return ov.api.send(ov)
}

// Send sends the photo.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what bots can send, please check the API documentation.
// On success, the sent message is returned as a MessageResponse.
func (op *OutgoingPhoto) Send() (*MessageResponse, error) {
	return op.api.send(op)
}

// Send sends the sticker.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what bots can send, please check the API documentation.
// On success, the sent message is returned as a MessageResponse.
func (os *OutgoingSticker) Send() (*MessageResponse, error) {
	return os.api.send(os)
}

// Send sends the audio.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what bots can send, please check the API documentation.
// On success, the sent message is returned as a MessageResponse.
func (oa *OutgoingAudio) Send() (*MessageResponse, error) {
	return oa.api.send(oa)
}

// Send sends the voice message.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what bots can send, please check the API documentation.
// On success, the sent message is returned as a MessageResponse.
func (ov *OutgoingVoice) Send() (*MessageResponse, error) {
	return ov.api.send(ov)
}

// Send sends the document.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what bots can send, please check the API documentation.
// On success, the sent message is returned as a MessageResponse.
func (od *OutgoingDocument) Send() (*MessageResponse, error) {
	return od.api.send(od)
}

// Send sends the request.
// On success, the photos are returned as a UserProfilePhotosResponse.
func (op *OutgoingUserProfilePhotosRequest) Send() (*UserProfilePhotosResponse, error) {
	resp := &UserProfilePhotosResponse{}
	_, err := op.api.c.postJSON(getUserProfilePhotos, resp, op)

	if err != nil {
		return nil, err
	}
	err = check(&resp.baseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Send sends the chat action.
// On success, nil is returned
func (oc *OutgoingChatAction) Send() error {
	resp := &baseResponse{}
	_, err := oc.api.c.postJSON(sendChatAction, resp, oc)

	if err != nil {
		return err
	}
	err = check(resp)
	if err != nil {
		return err
	}
	return nil
}

// Send sends the inline query answer.
// On success, nil is returned
func (ia *InlineQueryAnswer) Send() error {
	resp := &baseResponse{}
	_, err := ia.api.c.postJSON(answerInlineQuery, resp, ia)

	if err != nil {
		return err
	}
	err = check(resp)
	if err != nil {
		return err
	}
	return nil
}
