package user_service

func (this *UserController) LoginTest() {
	username := this.GetString("username")
	this.GetLogger().Msg("this is a message with trace id")
	this.GetLogger().Msgf("username: %s try to login.", username)

}