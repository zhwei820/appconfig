package default_service

func (this *DefaultController) GetAllPublic() {
	this.GetLogger().Msg("log from other service!!!!!")
}
