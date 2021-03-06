package selectors

import "github.com/Gameye/gameye-sdk-go/models"

/*
TemplateItem holds information about a template
*/
type TemplateItem = models.TemplateQueryArgItem

/*
SelectTemplateList selects a list of templates.
@param templateState template state
*/
func SelectTemplateList(
	templateState *models.TemplateQueryState,
) (
	templateList []*TemplateItem,
) {
	templateList = make([]*TemplateItem, 0)

	for _, templateItem := range templateState.Template {
		if templateItem == nil {
			continue
		}

		templateList = append(templateList, templateItem)
	}

	return
}

/*
SelectTemplateItem gets details about a single template from a
template-state as returned by the gameye api.
@param templateState template state
@param templateKey identifier of the template
*/
func SelectTemplateItem(
	templateState *models.TemplateQueryState,
	templateKey string,
) (
	templateItem *TemplateItem,
) {
	templateItem = templateState.Template[templateKey]
	return
}
