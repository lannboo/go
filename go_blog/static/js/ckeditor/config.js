/**
 * @license Copyright (c) 2003-2019, CKSource - Frederico Knabben. All rights reserved.
 * For licensing, see https://ckeditor.com/legal/ckeditor-oss-license
 */

CKEDITOR.editorConfig = function( config ) {
	// Define changes to default configuration here. For example:
	// config.language = 'fr';
	// config.uiColor = '#AADC6E';
    //
    // config.language ='zh-cn';
    // config.image_previewText = ' ';
    // config.removeDialogTabs = 'image:advanced;image:Link';//隐藏超链接与高级选项
    // config.filebrowserImageUploadUrl = "http://localhost/test";//上传图片的地址
    //



    //工具栏的配置

    config.toolbar = [
        ['Paste'],
        ["Bold","Italic","Underline"],
        ["NumberedLisst","BulletedList","-","JustifyLeft","JustifyCenter","JustifyRight","JustfyBlock"],
        ["Image"],
        ["Styles","Format","Font","FontSize"],
        ["TextColor","BGColor"],
        ["Maximize"]
       ]; //配置文件上传方法的路径 config.filebrowserUploadUrl="control/news/uploadImage.action"; var pathName = window.document.location.pathname; //获取带"/"的项目名，如：/uimcardprj var projectName = pathName.substring(0, pathName.substr(1).indexOf('/') + 1); config.filebrowserImageUploadUrl = projectName+'/control/news/uploadImage.action';



};
